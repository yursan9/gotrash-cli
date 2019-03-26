package lib

import "fmt"

func EmptyTrash() {
	n := len(getFilesInDir(trashInfoDir))
	fmt.Printf("Deleting %d files...\n", n)

	if prompt("Delete all trash permanently") {
		emptyDir(trashInfoDir)
		emptyDir(trashFilesDir)
	}
}
