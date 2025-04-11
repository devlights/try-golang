package main

import (
	"errors"
	"log"
	"net"
	"time"

	"golang.org/x/sys/unix"
)

func main() {
	log.SetFlags(log.Lmicroseconds)

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var (
		udsConn net.Conn
		err     error
	)
	for range 3 {
		udsConn, err = net.DialTimeout("unix", "@tcp_fd_passing", 1*time.Second)
		if err != nil {
			var netErr net.Error
			if errors.As(err, &netErr); netErr.Timeout() {
				continue
			}

			return err
		}

		break
	}
	defer udsConn.Close()
	log.Println("[TCP-S] connect uds-server")

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		return err
	}
	defer ln.Close()
	log.Println("[TCP-S] listen on :8888")

	for {
		errCh := make(chan error, 1)
		func() {
			conn, err := ln.Accept()
			if err != nil {
				errCh <- err
				return
			}
			defer func() {
				conn.Close()
				log.Println("[TCP-S] close")
			}()
			log.Println("[TCP-S] accept client")

			unixConn, _ := udsConn.(*net.UnixConn)
			tcpConn, _ := conn.(*net.TCPConn)
			file, _ := tcpConn.File()
			err = sendFD(unixConn, int(file.Fd()))
			if err != nil {
				errCh <- err
				return
			}
			log.Printf("[TCP-S] send fd=%d to uds-server", file.Fd())

			errCh <- nil
		}()

		err = <-errCh
		if err != nil {
			return err
		}
	}
}

func sendFD(sock *net.UnixConn, fd int) error {
	var (
		dummy  = make([]byte, 1)
		rights = unix.UnixRights(fd)
		err    error
	)
	_, _, err = sock.WriteMsgUnix(dummy, rights, nil)
	if err != nil {
		return err
	}

	return nil
}
