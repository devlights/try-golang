// server.go
// サーキットブレーカーの動作確認用 TCP サーバー。
// -err-rate で意図的にエラーを発生させる。
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"strings"
	"time"
)

func main() {
	log.SetFlags(log.Lmicroseconds)

	addr := flag.String("addr", ":9000", "listen address")
	errRate := flag.Float64("err-rate", 0.7, "error rate (0.0-1.0)")
	flag.Parse()

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
	defer ln.Close()
	log.Printf("[server] listening on %s  error_rate=%.0f%%", *addr, *errRate*100)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("[server] accept error: %v", err)
			continue
		}
		go handleConn(conn, *errRate)
	}
}

func handleConn(conn net.Conn, errRate float64) {
	defer conn.Close()
	remote := conn.RemoteAddr().String()
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		time.Sleep(50 * time.Millisecond) // 疑似処理遅延

		if rand.Float64() < errRate {
			resp := "ERR service_unavailable\n"
			fmt.Fprint(conn, resp)
			log.Printf("[server] %s → %s", remote, strings.TrimSpace(resp))
		} else {
			resp := fmt.Sprintf("OK hello_from_server req=%s\n", line)
			fmt.Fprint(conn, resp)
			log.Printf("[server] %s → %s", remote, strings.TrimSpace(resp))
		}
	}

	if sc.Err() != nil {
		log.Printf("sc.Err(): %v", sc.Err())
	}
}
