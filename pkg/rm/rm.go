package rm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"trash-cli/pkg/fs"
	"trash-cli/pkg/prompt"
	"trash-cli/pkg/trashinfo"
)

func Run(files []string, pattern string, interactive bool) {
	trashInfoDir := fs.GetTrashInfoDir()
	trashlist := trashinfo.NewTrashList(trashInfoDir)

	if pattern != "" {
		matched := trashlist.MatchTrash(pattern)
		files = append(files, matched...)
	}

	if interactive {
		for _, file := range files {
			fmt.Printf("Deleting... %s\n", file)
			if prompt.YesNo("Are you sure") {
				rmTrash(trashlist, file)
			}
		}
		return
	}

	if len(files) == 0 {
		fmt.Println("Trash is empty")
		os.Exit(0)
	}
	for _, file := range files {
		fmt.Printf("Deleting... %s\n", file)
	}

	if prompt.YesNo("Are you sure") {
		for _, file := range files {
			rmTrash(trashlist, file)
		}
	}
}

func rmTrash(trashlist trashinfo.TrashList, path string) {
	trashFilesDir := fs.GetTrashFilesDir()
	trashInfoFile := trashlist.FindTrash(path)

	if trashInfoFile == "" {
		fmt.Println("No such file in trash")
		os.Exit(1)
	}

	trashFile := strings.Replace(trashInfoFile, ".trashinfo", "", 1)
	trashFile = filepath.Base(trashFile)
	trashFile = filepath.Join(trashFilesDir, trashFile)

	fs.RmFile(trashInfoFile)
	fs.RmFile(trashFile)
}
