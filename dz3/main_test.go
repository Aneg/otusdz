package main

import (
	"testing"
)

func TestCalculateAndSort(t *testing.T) {
	testCases := struct {
		input  string
		output []Word
	}{
		input: "qw qw UJr er er qw UJr UJr t t qw t t t t ss t t ss t t ss t t t t t t ,df ss df df ss df df DS df AA GFD GFD",
		output: []Word{
			{str: "t", count: 16},
			{str: "df", count: 6},
			{str: "ss", count: 5},
			{str: "qw", count: 4},
			{str: "ujr", count: 3},
		},
	}

	words := CalculateAndSort([]byte(testCases.input), 5)

	if len(words) != len(testCases.output) {
		t.Errorf("Не вероное количество эддементов - %d", len(words))
	}

	for i, wordTest := range testCases.output {
		if wordTest.str != words[i].str || wordTest.count != words[i].count {
			t.Errorf("Не верный результат - %d", i)
		}
	}
}
