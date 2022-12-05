// Go でのソケットプログラミング 基本 (2)
//
// 本パッケージはサーバ側の処理です。
//
// # REFERENCES
//   - https://pkg.go.dev/net@go1.19.3
//   - https://www.developer.com/languages/intro-socket-programming-go/
//   - https://stackoverflow.com/questions/13417095/how-do-i-stop-a-listening-server-in-go
//   - https://stackoverflow.com/a/237495
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
)

const (
	PORT = 8888
)

var (
	appLog = log.New(os.Stderr, "[server] ", log.Ltime|log.Lmicroseconds)
)

func exitOnErr(err error) {
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
	}
}

func main() {
	// Start
	var (
		listener net.Listener
		err      error
	)

	listener, err = net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	exitOnErr(err)
	appLog.Println("server start...")

	// Regist SIGINT(Ctrl-C) handler
	var (
		quitCh = make(chan os.Signal, 1)
	)
	signal.Notify(quitCh, os.Interrupt)

	go func() {
		<-quitCh
		listener.Close()
	}()

	// Accept
	var (
		count int
	)

	for {
		var (
			conn net.Conn
		)

		conn, err = listener.Accept()
		if err != nil {
			appLog.Println("shutdown...")
			break
		}

		appLog.Printf("accept from %v", conn.RemoteAddr())
		go proc(conn)
		count++
	}

	log.Printf("COUNT=%d\n", count)
}

func proc(conn net.Conn) {
	defer conn.Close()

	// Recv
	//
	// Protocol:
	// 		(1) length: uint32 (4-bytes)
	// 		(2) data  : string (variable)
	var (
		totalBuf = new(bytes.Buffer)
		err      error
	)

	// (1) length
	var (
		buf    = make([]byte, 4)
		length uint32
	)

	_, err = conn.Read(buf)
	exitOnErr(err)

	length = binary.BigEndian.Uint32(buf[:])
	_, err = totalBuf.Write(buf)
	exitOnErr(err)

	// (2) data
	var (
		message string
	)

	buf = make([]byte, length)
	_, err = conn.Read(buf)
	exitOnErr(err)

	message = string(buf)
	_, err = totalBuf.Write(buf)
	exitOnErr(err)

	// Send
	var (
		resp    = strings.ToUpper(message)
		respBuf = []byte(resp)
	)
	_, err = conn.Write(respBuf)
	exitOnErr(err)

	appLog.Printf("\t[bytes] %v\n", totalBuf.Bytes())
}
