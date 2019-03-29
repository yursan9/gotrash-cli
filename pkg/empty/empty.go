package lib

import (
	"fmt"
	"os"
)

func EmptyTrash() {
	n := len(getFilesInDir(trashInfoDir))
	if n == 0 {
		fmt.Println("Trash is empty")
		os.Exit(0)
	}

	fmt.Printf("Deleting %d files...\n", n)
	if prompt("Delete all trash permanently") {
		emptyDir(trashInfoDir)
		emptyDir(trashFilesDir)
	}
}
