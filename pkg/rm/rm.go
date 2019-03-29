package rm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"trash-cli/pkg/fs"
	"trash-cli/pkg/trashinfo"
)

func Run(files []string) {
	for _, file := range files {
		fmt.Printf("Deleting... %s\n", file)
	}

	if prompt("Are you sure") {
		for _, file := range files {
			rmTrash(file)
		}
	}
}

func rmTrash(path string) {
	trashInfoDir := fs.GetTrashInfoDir()
	trashFilesDir := fs.GetTrashFilesDir()
	trashInfoList := trashinfo.NewTrashList(trashInfoDir)

	var trashInfoFile string
	for _, item := range trashInfoList {
		if item.Path == path {
			trashInfoFile = item.Name
			break
		}
	}

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

func prompt(text string) bool {
	var answer string

	fmt.Print(text, " [Yes/No]? ")
	fmt.Scanln(&answer)

	if answer == "y" || answer == "Y" || answer == "yes" || answer == "Yes" {
		return true
	}

	return false
}
