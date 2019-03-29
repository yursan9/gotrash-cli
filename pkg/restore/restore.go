package lib

import (
	"fmt"
	"path/filepath"
	"strings"
)

func RestoreTrash() {
	files := getFilesInDir(trashInfoDir)

	trashInfoList := newTrashInfoList(files)

	for i, item := range trashInfoList {
		var formattedTime = item.DeletionDate.Format("2006-01-02 15:04:05")
		fmt.Println(i+1, formattedTime, item.Path)
	}

	n := promptInt("Select number to restore") - 1
	t := trashInfoList[n]

	orig := filepath.Base(t.Name)
	orig = strings.Replace(orig, ".trashinfo", "", 1)
	orig = filepath.Join(trashFilesDir, orig)
	dest := t.Path

	moveFile(orig, dest)
	rmFile(t.Name)
}
