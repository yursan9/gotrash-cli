package restore

import (
	"fmt"
	"path/filepath"
	"strings"
	"trash-cli/pkg/fs"
	"trash-cli/pkg/trashinfo"
)

func Run() {
	trashInfoDir := fs.GetTrashInfoDir()
	trashFilesDir := fs.GetTrashFilesDir()
	trashInfoList := trashinfo.NewTrashList(trashInfoDir)

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

	fs.MoveFile(orig, dest)
	fs.RmFile(t.Name)
}

func promptInt(text string) int {
	var answer int

	fmt.Print(text, "? ")
	fmt.Scanln(&answer)

	return answer
}
