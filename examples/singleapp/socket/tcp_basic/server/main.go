// Go でのソケットプログラミング 基本 (1)
//
// 本パッケージはサーバ側の処理です。
//
// REFERENCES
//   - https://www.developer.com/languages/intro-socket-programming-go/
//   - https://stackoverflow.com/questions/13417095/how-do-i-stop-a-listening-server-in-go
//   - https://stackoverflow.com/a/237495
//

package main

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/devlights/gomy/errs"
)

var (
	appLog = log.New(os.Stdout, "[server]", 0)
	errLog = log.New(os.Stderr, "[server]", 0)
)

func main() {
	//
	// 起動
	//
	var (
		getAddr = net.ResolveTCPAddr
		laddr   = errs.Forget(getAddr("tcp", "localhost:8888"))
		server  *net.TCPListener
		err     error
	)

	server, err = net.ListenTCP("tcp", laddr)
	if err != nil {
		errLog.Fatal(err)
	}

	//
	// Ctrl-C や kill コマンドで停止要求が行われた場合のハンドリング
	//
	var (
		done  = make(chan struct{})
		sigCh = make(chan os.Signal, 1)
	)
	defer close(sigCh)

	go func() {
		<-sigCh
		appLog.Println("shutting down...")
		close(done)
		server.Close()
	}()

	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	//
	// 要求受付
	//
	appLog.Printf("Listening on %s\n", server.Addr().String())
	for {
		var (
			conn *net.TCPConn
		)

		conn, err = server.AcceptTCP()
		if err != nil {
			select {
			case <-done:
			default:
				errLog.Printf("error at AcceptTCP (%v)%T", err, err)
			}

			break
		}

		appLog.Printf("connected from: %s\n", conn.RemoteAddr())

		go func() {
			defer conn.Close()

			//
			// 受信
			//
			var (
				buf       = new(bytes.Buffer)
				bytesRead int
			)

			for {
				var (
					chunk     = make([]byte, 1)
					chunkSize int
				)

				chunkSize, err = conn.Read(chunk)
				appLog.Printf("chunk recv: %dbyte(s)\terr:%v", chunkSize, err)

				if err != nil {
					if errors.Is(err, io.EOF) {
						bytesRead += chunkSize
						buf.Write(chunk)
						break
					}

					appLog.Printf("error at conn.Read (%v)", err)
					return
				}

				if chunkSize == 0 {
					appLog.Printf("closed by remote")
					return
				}

				bytesRead += chunkSize
				buf.Write(chunk)
			}
			appLog.Printf("%d bytes recv", bytesRead)

			var (
				data = buf.Bytes()[:bytesRead]
			)
			appLog.Printf("recv: %s\n", data)

			//
			// 送信
			//
			var (
				message = bytes.ToUpper(data)
			)

			_, err = conn.Write(message)
			if err != nil {
				errLog.Printf("error at client writeto (%v)", err)
				return
			}

			err = conn.CloseWrite()
			if err != nil {
				errLog.Printf("error at conn closewrite (%v)", err)
				return
			}
		}()
	}
}
