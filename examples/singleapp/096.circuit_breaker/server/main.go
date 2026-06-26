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

	var (
		addr    = flag.String("addr", ":9000", "listen address")
		errRate = flag.Float64("err-rate", 0.7, "error rate (0.0-1.0)")
	)
	flag.Parse()

	var (
		ln  net.Listener
		err error
	)
	if ln, err = net.Listen("tcp", *addr); err != nil {
		log.Fatalf("listen error: %v", err)
	}
	defer ln.Close()
	log.Printf("[server] listening on %s  error_rate=%.0f%%", *addr, *errRate*100)

	for {
		var (
			conn net.Conn
		)
		if conn, err = ln.Accept(); err != nil {
			log.Printf("[server] accept error: %v", err)
			continue
		}

		go handleConn(conn, *errRate)
	}
}

func handleConn(conn net.Conn, errRate float64) {
	defer conn.Close()

	var (
		remote = conn.RemoteAddr().String()
		sc     = bufio.NewScanner(conn)
		line   string
	)
	for sc.Scan() {
		line = strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		time.Sleep(50 * time.Millisecond) // 擬似的な遅延

		var (
			resp string
		)
		if rand.Float64() < errRate {
			resp = "ERR service_unavailable\n"
			fmt.Fprint(conn, resp)
			log.Printf("[server] %s → %s", remote, strings.TrimSpace(resp))
		} else {
			resp = fmt.Sprintf("OK hello_from_server req=%s\n", line)
			fmt.Fprint(conn, resp)
			log.Printf("[server] %s → %s", remote, strings.TrimSpace(resp))
		}
	}

	if sc.Err() != nil {
		log.Printf("sc.Err(): %v", sc.Err())
	}
}
