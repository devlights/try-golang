// Go でのソケットプログラミング 基本 (2)
//
// 本パッケージはクライアント側の処理です。
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
	"math/rand"
	"net"
	"os"
	"time"
)

const (
	PORT = 8888
)

var (
	items = []string{
		"golang",
		"java",
		"python3",
		"C",
		"C++",
		"C#",
		"rust",
		"javascript",
	}
	appLog = log.New(os.Stderr, "[client] ", log.Ltime|log.Lmicroseconds)
)

func exitOnErr(err error) {
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
	}
}

func main() {
	// Connect
	var (
		conn net.Conn
		addr *net.TCPAddr
		err  error
	)

	addr, err = net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", PORT))
	exitOnErr(err)
	appLog.Printf("connect to %v", addr)

	conn, err = net.Dial("tcp", addr.String())
	exitOnErr(err)
	defer conn.Close()

	// Send
	//
	// Protocol:
	// 		(1) length: uint32 (4-bytes)
	// 		(2) data  : string (variable)
	var (
		rnd     = rand.New(rand.NewSource(time.Now().UnixNano()))
		message = items[rnd.Intn(len(items))]
		length  = make([]byte, 4)
		buf     = new(bytes.Buffer)
	)

	binary.BigEndian.PutUint32(length, uint32(len(message)))

	_, err = buf.Write(length)
	exitOnErr(err)
	_, err = buf.Write([]byte(message))
	exitOnErr(err)

	appLog.Printf("send %v --> %v", conn.LocalAddr(), conn.RemoteAddr())
	_, err = conn.Write(buf.Bytes())
	exitOnErr(err)

	// Recv
	var (
		resp []byte
	)

	resp, err = io.ReadAll(conn)
	exitOnErr(err)

	appLog.Printf("recv %v --> %v", conn.RemoteAddr(), conn.LocalAddr())
	appLog.Printf("\t%v --> %v", message, string(resp))
}
