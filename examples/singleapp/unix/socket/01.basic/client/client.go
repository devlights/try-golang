//go:build linux

package main

import (
	"log"
	"net"

	"golang.org/x/sys/unix"
)

func init() {
	log.SetFlags(log.Lmicroseconds)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		sfd int
		err error
	)

	sfd, err = unix.Socket(unix.AF_INET, unix.SOCK_STREAM, unix.IPPROTO_TCP)
	if err != nil {
		return err
	}
	defer func() {
		log.Println("[CLIENT] ソケットクローズ")
		unix.Close(sfd)
	}()

	var (
		ip   = net.ParseIP("127.0.0.1")
		ipv4 [4]byte

		sAddr unix.Sockaddr
	)
	copy(ipv4[:], ip.To4())

	sAddr = &unix.SockaddrInet4{Port: 8888, Addr: ipv4}
	err = unix.Connect(sfd, sAddr)
	if err != nil {
		return err
	}

	//
	// Send
	//
	var (
		buf = make([]byte, 2048)
		msg = "helloworld"
	)
	copy(buf, []byte(msg))

	err = unix.Send(sfd, buf[:len(msg)], 0)
	if err != nil {
		return err
	}

	//
	// Recv
	//
	var (
		n int
	)
	clear(buf)

	n, err = unix.Read(sfd, buf)
	if err != nil {
		return err
	}

	log.Printf("[CLIENT] %s", buf[:n])

	return nil
}
