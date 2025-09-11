package main

import (
	"bytes"
	_ "embed"
	"io"
	"log"
	"net"
	"os"

	"github.com/devlights/gomy/errs"
)

var (
	appLog = log.New(os.Stdout, "[server] ", 0)
	//go:embed main.go
	data []byte
)

func main() {
	//
	// Start
	//
	server, _ := net.ListenTCP("tcp", errs.Drop(net.ResolveTCPAddr("tcp", "localhost:8888")))
	defer func() {
		server.Close()
		appLog.Println("shutting down...")
	}()

	//
	// Accept
	//
	conn, _ := server.AcceptTCP()
	defer func() {
		appLog.Println("close connection...")
		conn.Close()
	}()

	//
	// Recv
	//
	buf := new(bytes.Buffer)
	io.Copy(buf, conn)
	appLog.Printf("%dbyte(s) recv", len(buf.Bytes()))

	//
	// Send
	//
	conn.Write(data)
	appLog.Printf("%dbyte(s) send", len(data))
	conn.CloseWrite()
	appLog.Println("close server-side write stream")
}
