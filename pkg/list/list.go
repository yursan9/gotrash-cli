package list

import (
	"fmt"
	"trash-cli/pkg/fs"
	"trash-cli/pkg/trashinfo"
)

func Run(terse bool) {
	trashDir := fs.GetTrashInfoDir()
	list := trashinfo.NewTrashList(trashDir)
	list.SortByDate()

	for _, item := range list {
		if terse {
			fmt.Println(item.Path)
			continue
		}
		var formattedTime = item.DeletionDate.Format("2006-01-02 15:04:05")
		fmt.Println(formattedTime, item.Path)
	}
}
