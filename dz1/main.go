package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
)

func main() {
	l := log.New(os.Stderr, "dz1-", log.Ldate|log.Ltime)
	if time, err := ntp.Time("0.beevik-ntp.pool.ntp.org"); err != nil {
		l.Fatalln(err)
	} else {
		fmt.Println(time.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
	}
}
