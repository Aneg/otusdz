package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func readRoutine(ctx context.Context, conn net.Conn, cancelFunc context.CancelFunc) {
	defer cancelFunc()
	scanner := bufio.NewScanner(conn)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN")
				break OUTER
			}
			text := scanner.Text()
			log.Printf("From server: %s", text)
		}
	}
	log.Printf("Finished readRoutine")
}

func writeRoutine(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				break OUTER
			}
			str := scanner.Text()
			log.Printf("To server %v\n", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}

	}
	log.Printf("Finished writeRoutine")
}

func main() {
	var timeoutStr string
	var timeout time.Duration
	var err error
	flag.StringVar(&timeoutStr, "timeout", "", "timeout")
	flag.Parse()

	if len(timeoutStr) != 0 {
		if timeout, err = time.ParseDuration(timeoutStr); err != nil {
			log.Println(err)
		}
	}

	address, err := getAddress()
	if err != nil {
		log.Fatal(err)
	}

	dialer := &net.Dialer{}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)

	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	go readRoutine(ctx, conn, cancel)
	go writeRoutine(ctx, conn)

	if timeout != 0 {
		log.Printf("Соединение оборвётся через %s", timeoutStr)
		<-time.NewTimer(timeout).C
		cancel()
		conn.Close()
	} else {
		<-ctx.Done()
	}

}

func getAddress() (address string, err error) {
	if len(flag.Args()) < 1 {
		err = errors.New("не верное количество аргументов")
	} else if addr := net.ParseIP(flag.Args()[0]); addr == nil {
		address = flag.Args()[0]
	} else if len(flag.Args()) < 2 {
		err = errors.New("не верное количество аргументов")
	} else {
		address = flag.Args()[0] + ":" + flag.Args()[1]
	}
	return address, err
}
