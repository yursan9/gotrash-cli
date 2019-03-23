package lib

import (
	"fmt"
)

func ListFiles() {
	files := getFilesInDir(trashInfoDir)

	trashInfoList := newTrashInfoList(files)

	for _, item := range trashInfoList {
		var formattedTime = item.DeletionDate.Format("2006-01-02 15:04:05")
		fmt.Println(formattedTime, item.Path)
	}
}
