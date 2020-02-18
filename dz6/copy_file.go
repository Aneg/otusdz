package dz6

import (
	"errors"
	"io"
	"os"
)

func CopyFile(fromPath string, toPath string, offset, limit int64) error {
	from, err := os.OpenFile(fromPath, os.O_RDWR, 0644)
	if err != nil {
		return errors.New("другие ошибки, например нет прав")
	}
	defer from.Close()

	to, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer to.Close()

	fi, err := from.Stat()
	if err != nil {
		return err
	}

	fromSize := fi.Size()
	if fromSize < offset {
		return errors.New("offset exceeds file size")
	}
	if limit == 0 || offset+limit > fromSize {
		limit = fromSize - offset
	}

	if _, err := from.Seek(offset, io.SeekStart); err != nil {
		return err
	}
	_, err = io.CopyN(to, from, limit)
	return err
}
