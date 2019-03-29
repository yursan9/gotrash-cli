package trashinfo

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
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
