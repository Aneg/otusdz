package dz6

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func CopyFile(from *os.File, to *os.File, offset, limit int64) error {
	fmt.Println(limit)
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
	fmt.Println(limit)
	var N int64
	N = 10000
	b := make([]byte, N)
	var currentCopy int64
	var lenRead int
	for currentCopy = 0; currentCopy < limit && (currentCopy+offset) < fromSize; {
		// есливдруг конец
		if currentCopy+N > limit {
			b = make([]byte, limit-currentCopy)
		}
		if lenRead, err = from.ReadAt(b, offset+currentCopy); err != nil {
			return err
		}
		if _, err = to.WriteAt(b, offset+currentCopy); err != nil {
			return err
		}
		currentCopy += int64(lenRead)
	}

	return err
}
