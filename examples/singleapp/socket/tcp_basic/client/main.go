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
	"errors"
	"io"
	"log"
	"net"
	"os"

	"github.com/devlights/gomy/errs"
)

var (
	appLog = log.New(os.Stdout, "[client]", 0)
	errLog = log.New(os.Stderr, "[client]", 0)
)

func main() {
	//
	// 接続
	//
	var (
		getAddr = net.ResolveTCPAddr
		laddr   = errs.Forget(getAddr("tcp", "localhost:"))
		raddr   = errs.Forget(getAddr("tcp", "localhost:8888"))
		conn    *net.TCPConn
		err     error
	)

	conn, err = net.DialTCP("tcp", laddr, raddr)
	if err != nil {
		errLog.Fatal(err)
	}
	defer conn.Close()

	appLog.Printf("connected to: %s\n", conn.RemoteAddr())

	//
	// 送信
	//
	var (
		message = []byte("hello world")
	)

	_, err = conn.Write(message)
	if err != nil {
		appLog.Printf("error at conn.Write (%v)", err)
		return
	}

	//
	// 受信
	//
	var (
		buf       = make([]byte, 4096)
		bytesRead int
	)

	bytesRead, err = conn.Read(buf)
	switch {
	case err != nil && !errors.Is(err, io.EOF):
		appLog.Printf("error at conn.Read (%v)", err)
		return
	case bytesRead == 0:
		appLog.Printf("closed by remote")
		return
	default:
		appLog.Printf("%d bytes recv", bytesRead)
	}

	appLog.Printf("recv: %s\n", buf[:bytesRead])
}
