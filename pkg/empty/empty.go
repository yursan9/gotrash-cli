package empty

import (
	"fmt"
	"os"
	"trash-cli/pkg/fs"
	"trash-cli/pkg/prompt"
)

func Run() {
	trashInfoDir := fs.GetTrashInfoDir()
	trashFilesDir := fs.GetTrashFilesDir()

	n := len(fs.GetFilesInDir(trashInfoDir))
	if n == 0 {
		fmt.Println("Trash is empty")
		os.Exit(0)
	}

	fmt.Printf("Deleting %d files...\n", n)
	if prompt.YesNo("Delete all trash permanently") {
		fs.EmptyDir(trashInfoDir)
		fs.EmptyDir(trashFilesDir)
	}
}
