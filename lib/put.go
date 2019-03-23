package lib

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/phayes/permbits"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func PutFiles(files []string) {
	for _, file := range files {
		PutFile(file)
	}
}

func PutFile(file string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("File didn't exist", file)
		return
	}

	absFile := ensureAbsPath(file)
	endFile := filepath.Base(file)

	parentDir := filepath.Dir(absFile)
	perm, _ := permbits.Stat(parentDir)

	if !perm.UserWrite() {
		fmt.Println("Don't have sufficent permissions to delete", file)
		os.Exit(1)
	}

	info := NewTrashInfo(absFile)
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

	err = moveFile(absFile, newPath)
	if err != nil {
		fmt.Println("Something wrong file move")
		os.Exit(1)
	}
}
