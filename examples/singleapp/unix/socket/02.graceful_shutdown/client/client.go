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
		_ = unix.Close(sfd)
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

	log.Println("[CLIENT] Connect")

	//
	// Send
	//
	var (
		buf = make([]byte, 2048)
		msg = "helloworld"
	)
	copy(buf, msg)

	err = unix.Send(sfd, buf[:len(msg)], 0)
	if err != nil {
		return err
	}

	log.Printf("[CLIENT] SEND %s", msg)

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

	log.Printf("[CLIENT] RECV %s", buf[:n])

	//
	// 正規解放 (Graceful Shutdown or Orderly Release)
	//
	// ソケットの正規解放とは、ソケット通信を適切に終了させ、リソースを解放するプロセスのことを指します。
	// これには通常、shutdownとcloseの2つの操作が含まれます。
	//
	// 1. Shutdown
	//   shutdownは通信相手に対して接続終了の意思を伝えます。
	//   例えば、SHUT_WRを使用すると、相手側にEOF（End of File）を送信します。
	//
	// 2. close
	//   closeはソケットのファイルディスクリプタを閉じ、関連するリソースを解放します。
	//   最後の参照が閉じられたときにのみ、ネットワークの端点を完全に解放します。
	//
	// 正規解放の手順
	//   1. shutdown(SHUT_WR) の呼び出し。これにより相手に送信停止を通知する。
	//   2. 必要に応じて、残りのデータを受信する。
	//   3. 最後に close を呼び出して、ソケットのリソースを完全に解放する。
	//
	// 正規解放を行うことで、ネットワーク通信を適切に終了し、リソースを効率的に管理することができます。
	// 特に信頼性の高い通信が必要な場合や、大規模なシステムでリソース管理が重要な場合に、この方法は有効です。
	//

	// 1. shutdown(SHUT_WR) の呼び出し。これにより相手に送信停止を通知する。
	//    つまり、相手側にEOFが送信される。「もうデータは送りません」という意思表示。
	err = unix.Shutdown(sfd, unix.SHUT_WR)
	if err != nil {
		return err
	}

	log.Println("[CLIENT] shutdown(SHUT_WR)")

	// 2. 必要に応じて、残りのデータを受信する。
LOOP:
	for {
		clear(buf)

		n, err = unix.Read(sfd, buf)
		switch {
		case n == 0:
			log.Println("[CLIENT] 切断検知 (0 byte read)")
			break LOOP
		case err != nil:
			if errors.Is(err, unix.ECONNRESET) {
				log.Printf("[CLIENT] 切断検知 (%s)", err)
				break LOOP
			}

			return err
		default:
			log.Printf("[CLIENT] RECV REMAIN [%s]", buf[:n])
		}
	}

	// 3. 最後に close を呼び出して、ソケットのリソースを完全に解放する。
	// これは上の defer で行われている。

	return nil
}
