package main

import (
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

type mapEl struct {
	Rune  rune
	Count int
}

func Decode(str string) string {
	strMap := make([]mapEl, 0, utf8.RuneCount([]byte(str))/2)
	for _, c := range []rune(str) {
		if num, err := strconv.Atoi(string(c)); err != nil {
			strMap = append(strMap, mapEl{Rune: c, Count: 0})
		} else {
			if len(strMap) == 0 {
				log.Fatalf("Некорректная строка")
			}
			strMap[len(strMap)-1].Count = strMap[len(strMap)-1].Count*10 + num
		}
	}

	var result strings.Builder
	for _, el := range strMap {
		result.WriteRune(el.Rune)
		for i := 1; i < el.Count; i++ {
			result.WriteRune(el.Rune)
		}
	}

	return result.String()
}
