package restore

import (
	"fmt"
	"path/filepath"
	"strings"
	"trash-cli/pkg/fs"
	"trash-cli/pkg/prompt"
	"trash-cli/pkg/trashinfo"
)

func Run() {
	trashInfoDir := fs.GetTrashInfoDir()
	trashFilesDir := fs.GetTrashFilesDir()
	trashInfoList := trashinfo.NewTrashList(trashInfoDir)

	for i, item := range trashInfoList {
		var formattedTime = item.DeletionDate.Format("2006-01-02 15:04:05")
		fmt.Println(i, formattedTime, item.Path)
	}

	n := prompt.Number("Select number to restore")
	t := trashInfoList[n]

	orig := filepath.Base(t.Name)
	orig = strings.Replace(orig, ".trashinfo", "", 1)
	orig = filepath.Join(trashFilesDir, orig)
	dest := t.Path

	fs.MoveFile(orig, dest)
	fs.RmFile(t.Name)
}
