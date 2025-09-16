//go:build linux

package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"net"
	"time"
)

type (
	Args struct {
		IsServer bool
	}
)

var (
	args Args
)

func init() {
	flag.BoolVar(&args.IsServer, "server", false, "server mode")
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	flag.Parse()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var err error
	switch args.IsServer {
	case true:
		err = runServer()
	default:
		err = runClient()
	}

	if err != nil {
		return err
	}

	return nil
}

func runServer() error {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		return err
	}
	defer ln.Close()

	errCh := make(chan error)
	defer close(errCh)

	for {
		select {
		case e := <-errCh:
			return e
		default:
		}

		err = ln.(*net.TCPListener).SetDeadline(time.Now().Add(1 * time.Second))
		if err != nil {
			return err
		}

		conn, err := ln.Accept()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}

			return err
		}

		go func(conn net.Conn) {
			defer func() {
				conn.Close()
				log.Println("[S] close")
			}()

			time.Sleep(100 * time.Millisecond)
			{
				log.Println("[S] send data")
				if _, err := conn.Write([]byte("hello")); err != nil {
					errCh <- err
				}

				buf := make([]byte, 1)
				for {
					_, err := conn.Read(buf)
					if err != nil {
						if errors.Is(err, io.EOF) {
							log.Println("[S] disconnect")
							break
						}

						errCh <- err
					}
				}
			}
		}(conn)
	}
}

func runClient() error {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		return err
	}
	defer func() {
		conn.Close()
		log.Println("[C] close")
	}()

	//
	// select(2) を使って読み込み可能かを判定
	// select(2)に指定するFDは、net.TCPConn.File() から取得する
	//
	tcpConn, _ := conn.(*net.TCPConn)
	file, err := tcpConn.File()
	if err != nil {
		return err
	}
	defer file.Close()

	var (
		fd   = SocketFd(file.Fd())
		buf  = make([]byte, 10)
		sec  = 0 * time.Second
		usec = 10 * time.Millisecond
	)
	for {
		readable, err := fd.Readable(sec, usec)
		if err != nil {
			return err
		}

		if !readable {
			log.Printf("[C] select(2) -- not readable(fd=%d)", int(fd))
			continue
		}

		clear(buf)
		n, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("[C] disconnect")
				break
			}

			return err
		}
		log.Printf("[C] recv %s", buf[:n])

		err = conn.(*net.TCPConn).CloseWrite()
		if err != nil {
			return err
		}
		log.Println("[C] shutdown(SHUT_WR)")

		break
	}

	return nil
}
