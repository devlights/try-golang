package main

import (
	"errors"
	"log"
	"net"
	"time"

	"github.com/devlights/fdpassing"
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
	log.Println("[TCP-S] tcp-listen on :8888")

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

			var (
				unixConn = udsConn.(*net.UnixConn)
				tcpConn  = conn.(*net.TCPConn)
				file, _  = tcpConn.File()
				fdp      = fdpassing.NewFd(unixConn)
			)
			err = fdp.Send(int(file.Fd()))
			if err != nil {
				errCh <- err
				return
			}
			log.Printf("[TCP-S] passing fd=%d to uds-server", file.Fd())

			errCh <- nil
		}()

		err = <-errCh
		if err != nil {
			return err
		}
	}
}
