/*
UNIXドメインソケット 抽象名前空間のサンプル（netパッケージのUnixConnを利用する版)
*/
package main

import (
	"bytes"
	"errors"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type (
	Args struct {
		IsServer bool
	}
)

const (
	// 抽象名前空間のソケットアドレス（@記号で始まる名前は\0に変換される）
	serverAddr = "@go_unix_domain_socket_test"
	bufSize    = len("hello")
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
	addr, err := net.ResolveUnixAddr("unix", serverAddr)
	if err != nil {
		return err
	}

	ln, err := net.ListenUnix("unix", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func(sigCh <-chan os.Signal) {
		<-sigCh
		ln.Close()
		os.Exit(0)
	}(sigCh)

	conn, err := ln.AcceptUnix()
	if err != nil {
		return err
	}
	defer conn.Close()

	buf := make([]byte, bufSize)
	for {
		clear(buf)

		n, err := conn.Read(buf)
		if n == 0 || errors.Is(err, io.EOF) {
			log.Println("[S] disconnect")
			break
		}

		if err != nil {
			if errors.Is(err, syscall.EINTR) {
				continue
			}
			return err
		}

		message := buf[:n]
		log.Printf("[S] Recv (%s)", message)

		message = bytes.ToUpper(buf[:n])
		_, err = conn.Write(message)
		if err != nil {
			return err
		}
		log.Printf("[S] Send (%s)", message)
	}

	return nil
}

func runClient() error {
	addr, err := net.ResolveUnixAddr("unix", serverAddr)
	if err != nil {
		return err
	}

	conn, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	buf := make([]byte, bufSize)
	copy(buf, []byte("hello"))

	_, err = conn.Write(buf)
	if err != nil {
		return err
	}
	log.Printf("[C] Send (%s)", buf)

	clear(buf)
	n, err := conn.Read(buf)
	if err != nil {
		if errors.Is(err, syscall.EINTR) {
			return nil
		}
		return err
	}

	if n == 0 {
		log.Println("[C] disconnect")
		return nil
	}

	message := string(buf[:n])
	log.Printf("[C] Recv (%s)", message)

	return nil
}
