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

	s := clearText(text)
	result := make(map[string]uint)
	textData := strings.Split(s, " ")
	log.Println(textData)
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

	count := 10
	if len(order) < 10 {
		count = len(order)
	}
	for key, _ := range order[:count] {
		log.Printf("%d) %s -> %d", key+1, order[key].str, order[key].count)
	}
}

func clearText(text []byte) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(string(text)), ",", ""), ".", "")
}

type Word struct {
	str   string
	count uint
}
