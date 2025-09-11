package main

import (
	"log"
	"net"
	"time"
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
	// 標準ライブライブラリで提供されている関数を使って実装
	//
	// 元のサンプルと同様にエラー処理は割愛する
	//
	var (
		ln net.Listener
	)
	ln, _ = net.Listen("tcp", ":8888")
	defer ln.Close()

	for {
		var (
			conn net.Conn
			addr *net.TCPAddr
		)
		conn, _ = ln.Accept()
		addr = conn.RemoteAddr().(*net.TCPAddr)

		log.Printf("[accept] EP: %v:%d", addr.IP, addr.Port)

		go func(conn net.Conn) {
			time.Sleep(1 * time.Second)
			conn.Close()
		}(conn)
	}
}
