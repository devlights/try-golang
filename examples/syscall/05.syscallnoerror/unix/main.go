//go:build unix

package main

import (
	"log"
	"net"
	"time"

	"golang.org/x/sys/unix"
)

func main() {
	log.SetFlags(0)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//
	// unix.SyscallNoError()を用いたサンプルと同じものを
	// unixパッケージで提供されている各ラッパー関数を使って実装
	//
	// 元のサンプルと同様にエラー処理は割愛する
	//

	// socket(2)
	var (
		fd int
	)
	fd, _ = unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	defer unix.Close(fd)

	// setsockopt(2) (SO_REUSEADDR)
	_ = unix.SetsockoptInt(fd, unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)

	// ソケットアドレス生成
	var (
		sAddr = unix.SockaddrInet4{
			Port: 8888,
			Addr: [4]byte{127, 0, 0, 1},
		}
	)

	// bind(2)
	_ = unix.Bind(fd, &sAddr)

	// listen(2)
	_ = unix.Listen(fd, 5)

	for {
		// accept(2)
		var (
			cfd int
			sa  unix.Sockaddr
			ca  *unix.SockaddrInet4
		)
		cfd, sa, _ = unix.Accept(fd)
		ca = sa.(*unix.SockaddrInet4)

		log.Printf("[accept] EP: %v:%d", net.IP(ca.Addr[:]), ca.Port)

		go func(fd int) {
			time.Sleep(1 * time.Second)
			_ = unix.Close(fd)
		}(cfd)
	}
}
