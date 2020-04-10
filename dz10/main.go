package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var timeoutStr string
	var timeout time.Duration
	var address string

	var err error
	flag.StringVar(&timeoutStr, "timeout", "", "timeout")
	flag.Parse()

	if timeout, err = getTimeout(timeoutStr); err != nil {
		log.Fatal(err)
	}
	if address, err = getAddress(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		log.Printf("Cannot connect: %v", err)
		cancel()
		return
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		readRoutine(ctx, conn, cancel, os.Stdin)
		wg.Done()
	}(&wg)
	go func(wg *sync.WaitGroup) {
		writeRoutine(ctx, conn, cancel, os.Stdin)
		wg.Done()
	}(&wg)

	select {
	case <-ctx.Done():
		conn.Write([]byte("Bye server!\n"))
	case <-c:
		log.Println("\r- Ctrl+C pressed in Terminal")
		cancel()
	}
	if err := conn.Close(); err != nil {
		log.Println(err)
	}
	wg.Wait()
}

func getTimeout(timeoutStr string) (timeout time.Duration, err error) {
	if len(timeoutStr) != 0 {
		if timeout, err = time.ParseDuration(timeoutStr); err != nil {
			return 0, err
		}
	} else {
		timeoutStr = "10s"
		timeout = time.Second * 10
	}
	return timeout, nil
}

func getAddress() (address string, err error) {
	if len(flag.Args()) < 1 {
		err = errors.New("неверное количество аргументов")
	} else if addr := net.ParseIP(flag.Args()[0]); addr == nil {
		address = flag.Args()[0]
	} else if len(flag.Args()) < 2 {
		err = errors.New("неверное количество аргументов")
	} else {
		address = flag.Args()[0] + ":" + flag.Args()[1]
	}
	return address, err
}

func readRoutine(ctx context.Context, conn net.Conn, cancelFunc context.CancelFunc, w io.Writer) {
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
			if _, err := w.Write([]byte(text)); err != nil {
				log.Println(err)
			}
			log.Printf("From server: %s", text)
		}
	}
	log.Printf("Finished readRoutine")
}

func writeRoutine(ctx context.Context, conn net.Conn, cancelFunc context.CancelFunc, stdin io.Reader) {
	scanner := bufio.NewScanner(stdin)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				log.Println("...EOF")
				cancelFunc()
				break OUTER
			}
			str := scanner.Text()
			if str == "quit" || str == "exit" {
				cancelFunc()
				break OUTER
			}
			log.Printf("To server %v\n", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}

	}
	log.Printf("Finished writeRoutine")
}
