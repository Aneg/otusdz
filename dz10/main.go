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
	"os/signal"
	"syscall"
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

func writeRoutine(ctx context.Context, conn net.Conn, cancelFunc context.CancelFunc) {
	scanner := bufio.NewScanner(os.Stdin)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				fmt.Println("...EOF")
				cancelFunc()
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
	} else {
		timeoutStr = "10s"
		timeout = time.Second * 10
	}

	address, err := getAddress()
	if err != nil {
		log.Fatal(err)
	}

	dialer := &net.Dialer{}
	ctx := context.Background()

	// Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		log.Printf("Cannot connect: %v", err)
		if timeout != 0 {
			//из дз: При подключении к несуществующему сервер, программа должна завершаться через timeout.
			log.Printf("Соединение оборвётся через %s", timeoutStr)
			<-time.NewTimer(timeout).C
		}
		cancel()
		return
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go readRoutine(ctx, conn, cancel)
	go writeRoutine(ctx, conn, cancel)

	select {
	case <-ctx.Done():
		conn.Write([]byte("Bye server!\n"))
	case <-c:
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		cancel()
	}
	if err := conn.Close(); err != nil {
		log.Println(err)
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
