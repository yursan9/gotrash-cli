package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RmTrashPrompted(files []string) {
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
	files := getFilesInDir(trashInfoDir)
	trashInfoList := newTrashInfoList(files)

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

	rmFile(trashInfoFile)
	rmFile(trashFile)
}
