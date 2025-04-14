package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/devlights/try-golang/examples/socket/tcp_fd_passing/fdpassing"
)

func main() {
	log.SetFlags(log.Lmicroseconds)

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ln, err := net.Listen("unix", "@tcp_fd_passing")
	if err != nil {
		return err
	}
	defer ln.Close()

	log.Println("[UDS-S] uds-server listening on")

	udsConn, err := ln.Accept()
	if err != nil {
		return err
	}
	defer udsConn.Close()

	log.Printf("[UDS-S] %v", udsConn.RemoteAddr())

	unixConn, ok := udsConn.(*net.UnixConn)
	if !ok {
		return fmt.Errorf("not net.UnixConn")
	}

	fd, err := fdpassing.NewFd(unixConn).Recv()
	if err != nil {
		return err
	}
	log.Printf("[UDS-S] recv fd=%d", fd)

	file := os.NewFile(uintptr(fd), "client-socket")
	if file == nil {
		return fmt.Errorf("os.NewFile() failed")
	}
	defer file.Close()

	conn, err := net.FileConn(file)
	if err != nil {
		return fmt.Errorf("net.FileConn() failed")
	}
	defer func() {
		conn.Close()
		log.Println("[UDS-S] close")
	}()

	buf := []byte("hello")
	_, err = conn.Write(buf)
	if err != nil {
		return err
	}
	log.Printf("[UDS-S] send (%s)", buf)

	buf = make([]byte, 5)
	n, err := conn.Read(buf)
	if err != nil {
		switch {
		case errors.Is(err, io.EOF):
			log.Println("[UDS-S] disconnect")
		default:
			return err
		}
	}
	log.Printf("[UDS-S] recv (%s)", buf[:n])

	tcpConn, _ := conn.(*net.TCPConn)
	tcpConn.CloseWrite()
	log.Println("[UDS-S] shutdown(SHUT_WR)")

	return nil
}
