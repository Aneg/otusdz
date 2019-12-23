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
		if test.output != Decode(test.input) {
			t.Errorf("Не верный результат тест %d", num)
		}
	}
}
