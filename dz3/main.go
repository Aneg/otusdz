package main

import (
	"io/ioutil"
	"log"
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

	r := make([]struct {
		str   string
		count uint
	}, len(result))

	for str, countStr := range result {
		lastKey := 0
		for i := len(r) - 1; i >= 0; i-- {
			if r[i].count >= countStr {
				lastKey = i
			} else {
				break
			}
		}
		r := append(r, struct {
			str   string
			count uint
		}{str: str, count: countStr})
	}

	log.Println(result)
}

func clearText(text []byte) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(string(text)), ",", ""), ".", "")
}
