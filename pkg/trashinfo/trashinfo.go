package trashinfo

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"
	"trash-cli/pkg/fs"
)

const TrashInfoTemp = `[Trash Info]
Path={{.Path | escape}}
DeletionDate={{.DeletionDate | format}}
`

type TrashInfo struct {
	Name         string
	Path         string
	DeletionDate time.Time
}

func New(path string) *TrashInfo {
	return &TrashInfo{
		Path:         path,
		DeletionDate: time.Now(),
	}
}

func (ti TrashInfo) WriteFile(path string) error {
	funcMap := template.FuncMap{
		"escape": func(uri string) string {
			path := url.PathEscape(uri)
			path = strings.Replace(path, "%2F", "/", -1)
			return path
		},
		"format": func(t time.Time) string {
			return t.Format("2006-01-02T15:04:05")
		},
	}
	t := template.Must(template.New("trashinfo").Funcs(funcMap).Parse(TrashInfoTemp))

	var content strings.Builder
	t.Execute(&content, ti)

	var filename strings.Builder
	fmt.Fprint(&filename, path, ".trashinfo")

	err := fs.AtomicWrite(filename.String(), content.String())
	if err != nil {
		return err
	}

	return nil
}

type TrashList []TrashInfo

func NewTrashList(dir string) TrashList {
	list := make(TrashList, 0)
	files := fs.GetFilesInDir(dir)

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		item := TrashInfo{
			Name: file,
		}
		for scanner.Scan() {
			if item.Path != "" && !item.DeletionDate.IsZero() {
				break
			}

			switch {
			case strings.HasPrefix(scanner.Text(), "Path="):
				path := strings.TrimPrefix(scanner.Text(), "Path=")
				item.Path, _ = url.PathUnescape(path)
			case strings.HasPrefix(scanner.Text(), "DeletionDate="):
				dateString := strings.TrimPrefix(scanner.Text(), "DeletionDate=")
				date, _ := time.Parse("2006-01-02T15:04:05", dateString)
				item.DeletionDate = date
			}
		}

		list = append(list, item)
	}

	return list
}

func (l TrashList) FindTrash(path string) string {
	l.SortByPath()

	i := sort.Search(len(l), func(i int) bool { return l[i].Path >= path })
	if i < len(l) && l[i].Path == path {
		return l[i].Name
	}

	return ""
}

func (l TrashList) MatchTrash(pattern string) []string {
	list := make([]string, 0)
	for _, item := range l {
		base := filepath.Base(item.Path)
		match, err := filepath.Match(pattern, base)
		if err != nil {
			fmt.Println("Bad pattern")
			os.Exit(1)
		}

		if match {
			list = append(list, item.Path)
		}
	}

	return list
}

func (l TrashList) SortByDate() {
	sort.Slice(l, func(i, j int) bool {
		return l[i].DeletionDate.Before(l[j].DeletionDate)
	})
}

func (l TrashList) SortByPath() {
	sort.Slice(l, func(i, j int) bool {
		return l[i].Path < l[j].Path
	})
}
