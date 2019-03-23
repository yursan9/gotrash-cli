package lib

import (
	"log"
	"os"
	"path/filepath"
)

var trashInfoDir string
var trashFilesDir string

func init() {
	xdgDataPath, ok := os.LookupEnv("XDG_DATA_HOME")
	if !ok {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Can't get Home directory")
		}
		xdgDataPath = filepath.Join(userHomeDir, ".local/share")
	}

	trashDir := filepath.Join(xdgDataPath, "Trash")

	trashInfoDir = filepath.Join(trashDir, "info")
	trashFilesDir = filepath.Join(trashDir, "files")
}

func getFilesInDir(path string) []string {
	files, _ := filepath.Glob(filepath.Join(path, "*"))
	return files
}

func atomicWrite(path, content string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(content)
	return nil
}

func moveFile(path, dest string) error {
	return os.Rename(path, dest)
}

func ensureDir(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func rmFile(path string) error {
	return os.Remove(path)
}

func emptyDir(path string) error {
	return os.RemoveAll(path)
}

func ensureAbsPath(path string) string {
	newPath, _ := filepath.Abs(path)
	return newPath
}
