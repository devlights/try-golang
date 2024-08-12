//go:build linux

package main

import (
	"errors"
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
	//
	// Create
	//
	var (
		sfd int
		err error
	)

	sfd, err = unix.Socket(unix.AF_INET, unix.SOCK_STREAM, unix.IPPROTO_TCP)
	if err != nil {
		return err
	}
	defer func() {
		log.Println("[SERVER] サーバーソケットクローズ")
		unix.Close(sfd)
	}()

	//
	// SO_REUSEADDR
	//
	err = unix.SetsockoptInt(sfd, unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
	if err != nil {
		return err
	}

	//
	// Bind and Listen
	//
	var (
		ip   = net.ParseIP("127.0.0.1")
		ipv4 [4]byte

		sAddr   unix.Sockaddr
		backLog = 2
	)
	copy(ipv4[:], ip.To4())

	sAddr = &unix.SockaddrInet4{Port: 8888, Addr: ipv4}
	err = unix.Bind(sfd, sAddr)
	if err != nil {
		return err
	}

	err = unix.Listen(sfd, backLog)
	if err != nil {
		return err
	}

	//
	// Accept
	//
	var (
		cfd   int
		cAddr unix.Sockaddr
	)

	cfd, cAddr, err = unix.Accept(sfd)
	if err != nil {
		return err
	}
	defer func() {
		log.Println("[SERVER] パケット送受信用ソケットクローズ")
		unix.Close(cfd)
	}()

	cAddrInet4 := cAddr.(*unix.SockaddrInet4)
	log.Printf("[SERVER] Connect from %v:%v", cAddrInet4.Addr, cAddrInet4.Port)

	//
	// Recv
	//
	var (
		buf = make([]byte, 2048)
		n   int
	)

	n, err = unix.Read(cfd, buf)
	if err != nil {
		return err
	}

	log.Printf("[SERVER] RECV %s", string(buf[:n]))

	//
	// Send
	//
	var (
		msg = "HELLOWORLD "
	)

	for range 5 {
		clear(buf)
		copy(buf, []byte(msg))

		err = unix.Send(cfd, buf[:len(msg)], 0)
		if err != nil {
			return err
		}

		log.Printf("[SERVER] SEND %s", buf[:len(msg)])
	}

	// 1. shutdown(SHUT_WR) の呼び出し。これにより相手に送信停止を通知する。
	err = unix.Shutdown(cfd, unix.SHUT_WR)
	if err != nil {
		return err
	}

	log.Println("[SERVER] shutdown(SHUT_WR)")

	// 2. 必要に応じて、残りのデータを受信する。
LOOP:
	for {
		clear(buf)

		n, err = unix.Read(cfd, buf)
		switch {
		case n == 0:
			log.Println("[SERVER] 切断検知 (0 byte read)")
			break LOOP
		case err != nil:
			var sysErr unix.Errno
			if errors.As(err, &sysErr); sysErr == unix.ECONNRESET {
				log.Printf("[SERVER] 切断検知 (%s)", sysErr)
				break LOOP
			}

			return err
		default:
			log.Printf("[SERVER] RECV %s", buf[:n])
		}
	}

	// 3. 最後に close を呼び出して、ソケットのリソースを完全に解放する。
	// これは上の defer で行われている。

	return nil
}
