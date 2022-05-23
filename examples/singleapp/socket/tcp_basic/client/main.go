// Go でのソケットプログラミング 基本 (1)
//
// 本パッケージはクライアント側の処理です。
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

	"github.com/devlights/gomy/errs"
)

var (
	appLog = log.New(os.Stdout, "[client] ", 0)
	errLog = log.New(os.Stderr, "[client] ", 0)
)

func connect(protocol, localAddr, remoteAddr string) (*net.TCPConn, error) {
	var (
		getAddr = net.ResolveTCPAddr
		laddr   = errs.Forget(getAddr(protocol, localAddr))
		raddr   = errs.Forget(getAddr(protocol, remoteAddr))
		conn    *net.TCPConn
		err     error
	)

	conn, err = net.DialTCP("tcp", laddr, raddr)
	if err != nil {
		return nil, err
	}

	return conn, err
}

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
			chunk     = make([]byte, 3)
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
	// 接続
	//
	conn, err := connect("tcp", "localhost:", "localhost:8888")
	if err != nil {
		errLog.Printf("error at connect (%v)", err)
	}
	defer conn.Close()

	appLog.Printf("connected to: %s\n", conn.RemoteAddr())

	//
	// 送信
	//
	var (
		message = []byte("hello world")
	)

	err = send(conn, message)
	if err != nil {
		errLog.Printf("error at conn.Write (%v)", err)
		return
	}
	appLog.Printf("%d bytes send", len(message))

	// 対向先にEOFを伝えるために無理やり送信側ストリームを閉じる
	// (本来は、通信メッセージ毎の構造規約があるはずなので、このようにすることは無い。サンプルなので。)
	err = conn.CloseWrite()
	if err != nil {
		errLog.Printf("error at conn.Write (%v)", err)
		return
	}
	appLog.Println("notify EOF (conn.CloseWrite)")

	//
	// 受信
	//
	data, err := recv(conn)
	if err != nil {
		errLog.Printf("error at recv (%v)", err)
	}
	appLog.Printf("%d bytes recv", len(data))
	appLog.Printf("recv: %s\n", data)
}
