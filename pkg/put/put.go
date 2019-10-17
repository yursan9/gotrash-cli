package put

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
	"trash-cli/pkg/fs"
	"trash-cli/pkg/trashinfo"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Run(files []string) {
	for _, file := range files {
		putFile(file)
	}
}

func putFile(file string) {
	trashInfoDir := fs.GetTrashInfoDir()
	trashFilesDir := fs.GetTrashFilesDir()

	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("File didn't exist", file)
		return
	}

	absFile := fs.EnsureAbsPath(file)
	endFile := filepath.Base(file)

	parentDir := filepath.Dir(absFile)

	if !fs.IsWriteable(parentDir) {
		fmt.Println("Don't have sufficent permissions to delete", file)
		os.Exit(1)
	}

	info := trashinfo.New(absFile)
	err := info.WriteFile(filepath.Join(trashInfoDir, endFile))

	var randNum int
	if os.IsExist(err) {
		var filename strings.Builder
		randNum = rand.Int()
		fmt.Fprintf(&filename, "%s-%d", endFile, randNum)
		err := info.WriteFile(filepath.Join(trashInfoDir, filename.String()))
		if err != nil {
			fmt.Println("Something wrong trashinfo")
			os.Exit(1)
		}
	}
	err = nil

	newPath := filepath.Join(trashFilesDir, endFile)
	if randNum != 0 {
		var filename strings.Builder
		fmt.Fprintf(&filename, "%s-%d", newPath, randNum)
		newPath = filename.String()
	}

	fs.MoveFile(absFile, newPath)
}
