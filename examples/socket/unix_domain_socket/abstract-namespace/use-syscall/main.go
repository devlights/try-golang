/*
UNIXドメインソケット 抽象名前空間のサンプル（syscallパッケージを利用する版)
*/
package main

import (
	"bytes"
	"errors"
	"flag"
	"log"
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
	serverAddr = "@go_unix_domain_socket_test"
	bufSize    = len("hello")
	backLog    = 1
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
	log.SetFlags(log.Lmicroseconds)

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
	fd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		return err
	}
	defer syscall.Close(fd)

	addr := &syscall.SockaddrUnix{
		Name: "\x00" + serverAddr[1:], // 抽象名前空間のソケットは、パス名の先頭にNULバイト(\0)を付けることで識別される
	}

	err = syscall.Bind(fd, addr)
	if err != nil {
		return err
	}

	err = syscall.Listen(fd, backLog)
	if err != nil {
		return err
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func(sigCh <-chan os.Signal) {
		<-sigCh
		syscall.Close(fd)
		os.Exit(0)
	}(sigCh)

	cfd, caddr, err := syscall.Accept(fd)
	if err != nil {
		return err
	}
	defer syscall.Close(cfd)

	unixAddr, ok := caddr.(*syscall.SockaddrUnix)
	if ok {
		// 「Connect from @」と出力されるが、これは正常な挙動。
		//   1.Unixドメインソケットの抽象名前空間では、クライアント側が明示的にバインドしていない場合、カーネルが自動的にアドレスを割り当てる
		//   2.この自動割り当てされたアドレスは通常表示されず、単に@または空の文字列として見えることがある.
		log.Printf("[S] Connect from %s", unixAddr.Name)
	}

	buf := make([]byte, bufSize)
	for {
		clear(buf)

		n, err := syscall.Read(cfd, buf)
		if err != nil {
			if errors.Is(err, syscall.EINTR) {
				continue
			}

			return err
		}

		if n == 0 {
			log.Println("[S] disconnect")
			break
		}

		message := buf[:n]
		log.Printf("[S] Recv (%s)", message)

		message = bytes.ToUpper(buf[:n])
		_, err = syscall.Write(cfd, message)
		if err != nil {
			return err
		}
		log.Printf("[S] Send (%s)", message)
	}

	return nil
}

func runClient() error {
	fd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		return err
	}
	defer syscall.Close(fd)

	addr := &syscall.SockaddrUnix{Name: "\x00" + serverAddr[1:]}
	err = syscall.Connect(fd, addr)
	if err != nil {
		return err
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func(sigCh <-chan os.Signal) {
		<-sigCh
		syscall.Close(fd)
		os.Exit(0)
	}(sigCh)

	buf := []byte("hello")
	_, err = syscall.Write(fd, buf)
	if err != nil {
		return err
	}
	log.Printf("[C] Send (%s)", string(buf))

	buf = make([]byte, bufSize)
	n, err := syscall.Read(fd, buf)
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
	log.Printf("[C] Recv (%s)", string(buf[:n]))

	return nil
}
