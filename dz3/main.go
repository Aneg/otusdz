package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	var text []byte
	var err error
	if text, err = ioutil.ReadFile("text.txt"); err != nil {
		log.Fatal(err)
	}

	top := CalculateAndSort(text, 10)
	for key := range top {
		log.Printf("%d) %s -> %d", key+1, top[key].str, top[key].count)
	}
}

func clearText(text []byte) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(string(text)), ",", ""), ".", "")
}

func CalculateAndSort(text []byte, topCount int) []Word {
	s := clearText(text)
	result := make(map[string]uint)
	textData := strings.Split(s, " ")
	for _, str := range textData {
		result[str] += 1
	}

	order := make([]Word, 0, len(result))

	for str, countStr := range result {
		order = append(order, Word{str: str, count: countStr})
	}

	sort.Slice(order, func(i, j int) bool {
		return order[i].count > order[j].count
	})

	if len(order) < topCount {
		topCount = len(order)
	}

	return order[:topCount]
}
