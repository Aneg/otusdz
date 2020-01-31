package dz6

import (
	"os"
	"testing"
)

func TestCopyFile(t *testing.T) {
	fromPath := "from.data"
	toPath := "to.data"

	from, err := os.OpenFile(fromPath, os.O_RDWR, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			t.Error("IsNotExist")
		}
		t.Error("другие ошибки, например нет прав")
	}
	defer from.Close()
	to, err := os.OpenFile(toPath, os.O_RDWR, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			t.Error("IsNotExist")
		}
		t.Error("другие ошибки, например нет прав")
	}
	defer to.Close()

	if err := CopyFile(from, to, 100, 100); err != nil {
		t.Error(err)
	}
}
