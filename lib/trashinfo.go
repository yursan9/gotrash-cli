package lib

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"sort"
	"strings"
	"text/template"
	"time"
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

func NewTrashInfo(path string) *TrashInfo {
	return &TrashInfo{
		Path:         path,
		DeletionDate: time.Now(),
	}
}

func (ti TrashInfo) WriteFile(path string) error {
	funcMap := template.FuncMap{
		"escape": func(uri string) string {
			return url.PathEscape(uri)
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

	err := atomicWrite(filename.String(), content.String())
	if err != nil {
		return err
	}

	return nil
}

func newTrashInfoList(files []string) []TrashInfo {
	list := make([]TrashInfo, 0)

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

	sort.Sort(ByDeletion(list))

	return list
}

type ByDeletion []TrashInfo

func (by ByDeletion) Len() int           { return len(by) }
func (by ByDeletion) Swap(i, j int)      { by[i], by[j] = by[j], by[i] }
func (by ByDeletion) Less(i, j int) bool { return by[i].DeletionDate.Before(by[j].DeletionDate) }
