package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"golang.org/x/sys/unix"
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

	log.Println("[UDS-S] server listening on")

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

	fd, err := recvFD(unixConn)
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

func recvFD(sock *net.UnixConn) (int, error) {
	var (
		dummy = make([]byte, 1)
		oob   = make([]byte, unix.CmsgSpace(4))
		flags int
		err   error
	)
	_, _, flags, _, err = sock.ReadMsgUnix(dummy, oob)
	if err != nil {
		return -1, err
	}

	if flags&unix.MSG_TRUNC != 0 {
		return -1, fmt.Errorf("control message is truncated")
	}

	var (
		msgs []unix.SocketControlMessage
	)
	msgs, err = unix.ParseSocketControlMessage(oob)
	if err != nil {
		return -1, err
	}

	if len(msgs) != 1 {
		return -1, fmt.Errorf("want: 1 control message; got: %d", len(msgs))
	}

	var (
		fds []int
	)
	fds, err = unix.ParseUnixRights(&msgs[0])
	if err != nil {
		return -1, err
	}

	if len(fds) != 1 {
		return -1, fmt.Errorf("want: 1 fd; got: %d", len(fds))
	}

	return fds[0], nil
}
