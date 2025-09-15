//go:build linux

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"syscall"
)

type (
	Args struct {
		UseListenConfig bool
	}
)

const (
	SO_REUSEADDR = syscall.SO_REUSEADDR
	SO_REUSEPORT = 0xf
)

var (
	args Args
)

func main() {
	log.SetFlags(log.Lmicroseconds)

	flag.BoolVar(&args.UseListenConfig, "listenconfig", false, "Use net.ListenConfig")
	flag.Parse()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var err error
	switch args.UseListenConfig {
	case true:
		err = runListenConfig()
	default:
		err = runNormal()
	}

	if err != nil {
		return err
	}

	return nil
}

func runListenConfig() error {
	lc := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			var (
				reuseAddr, reusePort int
				opErr1, opErr2       error
				err                  error
			)

			err = c.Control(func(fd uintptr) {
				reuseAddr, opErr1 = syscall.GetsockoptInt(int(fd), syscall.SOL_SOCKET, SO_REUSEADDR)
				reusePort, opErr2 = syscall.GetsockoptInt(int(fd), syscall.SOL_SOCKET, SO_REUSEPORT)
			})

			if err != nil {
				return err
			}

			log.Printf("SO_REUSEADDR=%d", reuseAddr)
			log.Printf("SO_REUSEPORT=%d", reusePort)

			if opErr1 != nil {
				return opErr1
			}

			if opErr2 != nil {
				return opErr2
			}

			return nil
		},
	}

	ln, err := lc.Listen(context.Background(), "tcp", ":8888")
	if err != nil {
		return err
	}
	defer ln.Close()

	return nil
}

func runNormal() error {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		return err
	}
	defer ln.Close()

	tcpLn, ok := ln.(*net.TCPListener)
	if !ok {
		return fmt.Errorf("not a TCP Listener")
	}

	file, err := tcpLn.File()
	if err != nil {
		return err
	}
	defer file.Close()

	var (
		fd = file.Fd()
		v  int
	)
	v, err = syscall.GetsockoptInt(int(fd), syscall.SOL_SOCKET, SO_REUSEADDR)
	if err != nil {
		return err
	}
	log.Printf("SO_REUSEADDR=%d", v)

	v, err = syscall.GetsockoptInt(int(fd), syscall.SOL_SOCKET, SO_REUSEPORT)
	if err != nil {
		return err
	}
	log.Printf("SO_REUSEPORT=%d", v)

	return nil
}
