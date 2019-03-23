package lib

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestAtomicWrite(t *testing.T) {
	dir, err := ioutil.TempDir("", "trash-cli")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dir)

	value := "Hello World!"
	expectedValue := []byte("Hello World!")

	fn := filepath.Join(dir, "example")
	atomicWrite(fn, value)

	content, err := ioutil.ReadFile(fn)
	if err != nil {
		t.Errorf("No test file")
	}

	if !bytes.Equal(expectedValue, content) {
		t.Errorf("Expect %s got %s", expectedValue, content)
	}
}

func TestMoveFile(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "a")
	if err != nil {
		t.Error(err)
	}
	tmpfile.Close()
	
	renamedtmp := filepath.Dir(tmpfile.Name())
	renamedtmp = filepath.Join(renamedtmp, "b")
	moveFile(tmpfile.Name(), renamedtmp)
	
	if _, err := os.Stat(renamedtmp); os.IsNotExist(err) {
		os.Remove(tmpfile.Name())
		t.Errorf("Can't rename %s to %s", tmpfile.Name(), renamedtmp)
	}
	os.Remove(renamedtmp)
}

func TestEnsureDir(t *testing.T) {
	tmpDir := os.TempDir()
	dir := filepath.Join(tmpDir, "test", "a", "dir")
	ensureDir(dir, 0700)
	
	if err := os.Chdir(dir); err != nil {
		t.Errorf("Can't make %s directory", dir)
	}
	os.RemoveAll(tmpDir)
}

func TestRmFile(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "a")
	if err != nil {
		t.Error(err)
	}
	tmpfile.Close()
	
	rmFile(tmpfile.Name())
	if _, err := os.Stat(tmpfile.Name()); os.IsExist(err) {
		t.Errorf("Can't delete %s file", tmpfile.Name())
	}
}
