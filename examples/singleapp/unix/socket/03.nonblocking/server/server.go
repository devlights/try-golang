//go:build linux

package main

import (
	"errors"
	"log"
	"net"
	"time"

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
	log.Println("[SERVER] set SO_REUSEADDR")

	//
	// Accept処理をノンブロッキングモードで処理するために設定
	//
	err = unix.SetNonblock(sfd, true)
	if err != nil {
		return err
	}
	log.Println("[SERVER] set O_NONBLOCK")

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
	//   ノンブロッキングモードにしているので
	//   接続するまでブロックされずに unix.EAGAIN が返る.
	//
	var (
		cfd   int
		cAddr unix.Sockaddr
	)

	for {
		cfd, cAddr, err = unix.Accept(sfd)
		if err != nil {
			if errors.Is(err, unix.EAGAIN) || errors.Is(err, unix.EWOULDBLOCK) || errors.Is(err, unix.EINTR) {
				log.Println("[SERVER][ACCEPT] --> unix.EAGAIN")

				time.Sleep(100 * time.Millisecond)
				continue
			}

			return err
		}

		break
	}

	defer func() {
		log.Println("[SERVER] パケット送受信用ソケットクローズ")
		unix.Close(cfd)
	}()

	cAddrInet4 := cAddr.(*unix.SockaddrInet4)
	log.Printf("[SERVER] Connect from %v:%v", cAddrInet4.Addr, cAddrInet4.Port)

	//
	// （補足）
	//    サーバソケットとAcceptで受け取ったパケット送受信用ソケットは別物なので
	//    受信と送信をノンブロッキングしたい場合は、再度 unix.SetNonblock(cfd, true) が必要.
	//
	// 本サンプルではブロッキングモードのままで処理している
	//

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

	// クライアントから受信した値を使って「何かの処理」を行った後に
	// クライアント側に返送するという流れをシミュレートするために
	// 意図的に少しディレイを入れる
	time.Sleep(150 * time.Millisecond)

	//
	// Send
	//
	var (
		msg = "HELLOWORLD "
	)

	for range 5 {
		clear(buf)
		copy(buf, []byte(msg))

		_, err = unix.Write(cfd, buf[:len(msg)])
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
