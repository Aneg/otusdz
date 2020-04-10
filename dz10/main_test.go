package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"testing"
	"time"
)

type ReaderWriter struct {
	ChRead chan []byte
}

func (rw *ReaderWriter) Write(p []byte) (n int, err error) {
	rw.ChRead <- p
	return 0, nil
}

//func TestReadRoutine(t *testing.T) {
//	message := "Hi there!\n"
//	ctx := context.Background()
//	ctx, cancel := context.WithTimeout(ctx, 3* time.Second)
//
//	readerWriter := &ReaderWriter{
//		ChRead: make(chan []byte, 10),
//	}
//
//	go func() {
//		conn, err := net.Dial("tcp", ":3000")
//		if err != nil {
//			t.Fatal(err)
//		}
//		defer conn.Close()
//
//		readRoutine(ctx, conn, cancel, readerWriter)
//	}()
//
//	l, err := net.Listen("tcp", ":3000")
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer l.Close()
//	for {
//		conn, err := l.Accept()
//		if err != nil {
//			return
//		}
//		defer conn.Close()
//
//		if _, err := fmt.Fprintf(conn, message); err != nil {
//			t.Fatal(err)
//		}
//
//		if r := string(<-readerWriter.ChRead); "Hi there!" != r {
//			t.Error("пришло не то сообщение ", r)
//		}
//		return // Done
//	}
//}

func TestWriteRoutine(t *testing.T) {
	message := "Hi there!\n"
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

	buffer := bytes.Buffer{}
	buffer.Write([]byte(message))
	go func() {
		conn, err := net.Dial("tcp", ":3000")
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		writeRoutine(ctx, conn, cancel, &buffer)
	}()

	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			t.Fatal(err)
		}

		if msg := string(buf[:]); msg != message {
			t.Fatalf("Unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", msg, message)
		}
		return // Done
	}

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
