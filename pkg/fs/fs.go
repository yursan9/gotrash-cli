package fs

import (
	"log"
	"os"
	"path/filepath"
)

func GetTrashDir() string {
	xdgDataPath, ok := os.LookupEnv("XDG_DATA_HOME")
	if !ok {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Can't get Home directory")
		}
		xdgDataPath = filepath.Join(userHomeDir, ".local/share")
	}

	trashDir := filepath.Join(xdgDataPath, "Trash")
	return trashDir
}

func GetTrashInfoDir(xdg string) string {
	trashDir := GetTrashDir()
	trashInfoDir = filepath.Join(trashDir, "info")

	ensureDir(trashInfoDir, 0700)
	
	return trashInfoDir
}

func GetTrashFilesDir(xdg string) string {
	trashDir := GetTrashDir()
	trashFilesDir = filepath.Join(trashDir, "files")

	ensureDir(trashFilesDir, 0700)
	
	return trashFilesDir
}

func GetFilesInDir(path string) []string {
	files, _ := filepath.Glob(filepath.Join(path, "*"))
	return files
}

func AtomicWrite(path, content string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(content)
	return nil
}

func MoveFile(path, dest string) error {
	return os.Rename(path, dest)
}

func EnsureDir(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func RmFile(path string) error {
	return os.Remove(path)
}

func EmptyDir(path string) error {
	return os.RemoveAll(path)
}

func EnsureAbsPath(path string) string {
	newPath, _ := filepath.Abs(path)
	return newPath
}
