package dz6

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCopyFile_with_limit_and_offset(t *testing.T) {
	fromPath := "from.data"
	toPath := "to.data"

	if err := createFile(fromPath, 1024*1024*10); err != nil {
		t.Error(err)
	}

	testCases := []struct {
		limit  int64
		offset int64
		first  int
		last   int
		length int
	}{
		{
			limit:  100,
			offset: 100,
			first:  49,
			last:   50,
			length: 100,
		},
		{
			limit:  0,
			offset: 100,
			first:  49,
			last:   50,
			length: 1024*1024*10 - 100,
		},
		{
			limit:  102,
			offset: 0,
			first:  49,
			last:   50,
			length: 102,
		},
		{
			limit:  0,
			offset: 0,
			first:  49,
			last:   50,
			length: 1024 * 1024 * 10,
		},
	}

	from, err := os.OpenFile(fromPath, os.O_RDWR, 0644)
	if err != nil {
		t.Error("другие ошибки, например нет прав")
	}
	defer from.Close()

	for num, testCase := range testCases {
		to, err := os.Create(toPath)
		if err != nil {
			t.Error(err)
		}
		if err := CopyFile("from.data", "to.data", testCase.offset, testCase.limit); err != nil {
			t.Errorf("№ %d: %s", num, err)
		}

		b, err := ioutil.ReadAll(to)
		if len(b) != int(testCase.length) {
			t.Errorf("№ %d: Не верно скопировано: %d", num, len(b))
		}
		if b[0] != byte(testCase.first) && b[testCase.length-1] != byte(testCase.last) {
			t.Errorf("№ %d: Не верно скопировано.", num)
		}

		to.Close()
		_ = os.Remove(toPath)
	}

	_ = os.Remove(fromPath)
}

func createFile(path string, size int) error {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}

	b := make([]byte, size)

	for i := 0; i < size; i++ {
		if i < 101 {
			b[i] = 49
		} else {
			b[i] = 50
		}
	}
	_, err = file.Write(b)
	return err
}
