package fs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/phayes/permbits"
)

// GetTrashDir returns the location of trash directory
// Usually it's $XDG_DATA_HOME/Trash or $HOME/.local/share/Trash
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

// GetTrashInfoDir returns the location of trash info directory
// that contains information about deleted files
// Usually it's TrashDir/info
func GetTrashInfoDir() string {
	trashDir := GetTrashDir()
	trashInfoDir := filepath.Join(trashDir, "info")

	EnsureDir(trashInfoDir, 0700)

	return trashInfoDir
}

// GetTrashFilesDir returns the location of trash' files directory
// that contains the deleted files
// Usually it's TrashDir/files
func GetTrashFilesDir() string {
	trashDir := GetTrashDir()
	trashFilesDir := filepath.Join(trashDir, "files")

	EnsureDir(trashFilesDir, 0700)

	return trashFilesDir
}

// GetFilesInDir returns all relative path of files in directory
func GetFilesInDir(path string) []string {
	files, _ := filepath.Glob(filepath.Join(path, "*"))
	return files
}

// AtomicWrite do an atomic write of content on file specify by path
func AtomicWrite(path, content string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(content)
	return nil
}

// IsWriteable checks if user have path write permission
func IsWriteable(path string) bool {
	perm, _ := permbits.Stat(path)
	return perm.UserWrite()
}

// IsExist checks if path is exist
func IsExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// MoveFile moves or renames path to dest
func MoveFile(path, dest string) error {
	return os.Rename(path, dest)
}

// EnsureDir makes directory on path with given permission
// Also make the parent directory if it not exist
func EnsureDir(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// RmFile delete path from filestystem
func RmFile(path string) error {
	return os.Remove(path)
}

// EmptyDir deletes all files on path
func EmptyDir(path string) error {
	return os.RemoveAll(path)
}

// EnsureAbsPath returns absolute path of file
func EnsureAbsPath(path string) string {
	newPath, _ := filepath.Abs(path)
	return newPath
}
