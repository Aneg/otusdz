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
	fmt.Printf("%s -> %s\n", str, Decode(str))
}
