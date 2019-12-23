package main

import (
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "w2r3fg4",
			output: "wwrrrfgggg",
		},
		{
			input:  "W5r11",
			output: "WWWWWrrrrrrrrrrr",
		},
	}

	for num, test := range tests {
		if r, err := Decode(test.input); err != nil || test.output != r {
			t.Errorf("Не верный результат тест %d", num)
		}
	}
}

func TestDecodeFail(t *testing.T) {
	if _, err := Decode("32"); err == nil {
		t.Error("Отсутствует ошибка")
	}
}
