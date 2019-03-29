package list

import (
	"fmt"
	"trash-cli/pkg/fs"
	"trash-cli/pkg/trashinfo"
)

func Run() {
	trashDir := fs.GetTrashInfoDir()
	list := trashinfo.NewTrashList(trashDir)

	for _, item := range list {
		var formattedTime = item.DeletionDate.Format("2006-01-02 15:04:05")
		fmt.Println(formattedTime, item.Path)
	}
}
