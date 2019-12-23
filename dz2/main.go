package main

import (
	"fmt"
	"log"
)

func main() {
	var str string
	if _, err := fmt.Scan(&str); err != nil {
		log.Fatalf("scan error: %s: ", err.Error())
	}
	var result string
	var err error
	if result, err = Decode(str); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s -> %s\n", str, result)
}
