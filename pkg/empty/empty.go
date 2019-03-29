package empty

import (
	"fmt"
	"os"
	"trash-cli/pkg/fs"
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
	if prompt("Delete all trash permanently") {
		fs.EmptyDir(trashInfoDir)
		fs.EmptyDir(trashFilesDir)
	}
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
