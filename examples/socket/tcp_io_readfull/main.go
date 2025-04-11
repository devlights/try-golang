package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

type (
	Args struct {
		bufsize int
		length  int
		timeout time.Duration
	}
)

var (
	args Args
)

func init() {
	flag.IntVar(&args.bufsize, "bufsize", 4, "bufsize")
	flag.IntVar(&args.length, "length", 2, "length")
	flag.DurationVar(&args.timeout, "timeout", 1*time.Second, "timeout")
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	flag.Parse()

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}
			panic(err)
		}
		defer conn.Close()

		buf := []byte(strings.Repeat("h", args.length))
		_, err = conn.Write(buf)
		if err != nil {
			panic(err)
		}

		time.Sleep(500 * time.Millisecond)
	}()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		return err
	}
	defer conn.Close()

	log.Println("[C] recv start")
	{
		buf := make([]byte, args.bufsize)
		for {
			clear(buf)

			err = conn.SetReadDeadline(time.Now().Add(args.timeout))
			if err != nil {
				return err
			}

			n, err := io.ReadFull(conn, buf)
			if err != nil {
				switch {
				case errors.Is(err, io.EOF):
					log.Println("io.EOF")
					return nil
				case errors.Is(err, io.ErrUnexpectedEOF):
					log.Println("io.ErrUnexpectedEOF")
					return nil
				default:
					var netErr net.Error
					if errors.As(err, &netErr) && netErr.Timeout() {
						log.Println("netErr.Timeout()")
						return nil
					}
				}

				return err
			}

			log.Printf("[C] data=(%s)", buf[:n])
		}
	}
}
