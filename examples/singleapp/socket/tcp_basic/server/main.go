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
	appLog = log.New(os.Stdout, "[server] ", 0)
	errLog = log.New(os.Stderr, "[server] ", 0)
)

func send(conn *net.TCPConn, b []byte) error {
	_, err := conn.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func recv(conn *net.TCPConn) ([]byte, error) {
	var (
		buf       = new(bytes.Buffer)
		bytesRead int
	)

	for {
		var (
			chunk     = make([]byte, 10)
			chunkSize int
			err       error
		)

		chunkSize, err = conn.Read(chunk)
		appLog.Printf("chunk recv: %dbyte(s)\terr:%v", chunkSize, err)

		if err != nil {
			if errors.Is(err, io.EOF) {
				bytesRead += chunkSize
				buf.Write(chunk)
				break
			}

			return nil, err
		}

		if chunkSize == 0 {
			return nil, errors.New("closed by remote")
		}

		bytesRead += chunkSize
		buf.Write(chunk)
	}

	return buf.Bytes()[:bytesRead], nil
}

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
			data, err := recv(conn)
			if err != nil {
				errLog.Printf("error at recv (%v)", err)
			}
			appLog.Printf("%d bytes recv", len(data))
			appLog.Printf("recv: %s\n", data)

			//
			// 送信
			//
			var (
				message = bytes.ToUpper(data)
			)

			err = send(conn, message)
			if err != nil {
				errLog.Printf("error at send (%v)", err)
				return
			}
			appLog.Printf("%d bytes send", len(message))

			// 対向先にEOFを伝えるために無理やり送信側ストリームを閉じる
			// (本来は、通信メッセージ毎の構造規約があるはずなので、このようにすることは無い。サンプルなので。)
			//
			// サーバ側の今回の実装では、これを行わなくても defer conn.Close() で切断するので対向先にEOFが通知されるが一応入れている			
			err = conn.CloseWrite()
			if err != nil {
				errLog.Printf("error at conn closewrite (%v)", err)
				return
			}
			appLog.Println("notify EOF (conn.CloseWrite)")


		}()
	}
}
