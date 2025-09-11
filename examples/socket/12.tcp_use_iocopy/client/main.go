package main

import (
	"bytes"
	_ "embed"
	"io"
	"log"
	"net"
	"os"
)

var (
	appLog = log.New(os.Stdout, "[client] ", 0)
	//go:embed main.go
	data []byte
)

func main() {
	//
	// Connect
	//
	laddr, _ := net.ResolveTCPAddr("tcp", "localhost:")
	raddr, _ := net.ResolveTCPAddr("tcp", "localhost:8888")

	conn, _ := net.DialTCP("tcp", laddr, raddr)
	defer func() {
		appLog.Println("close connection...")
		conn.Close()
	}()

	//
	// Send
	//
	conn.Write(data)
	appLog.Printf("%dbyte(s) send", len(data))

	conn.CloseWrite()
	appLog.Println("close client-side write stream")

	//
	// Recv
	//
	buf := new(bytes.Buffer)
	io.Copy(buf, conn)
	appLog.Printf("%dbyte(s) recv", len(buf.Bytes()))
}
