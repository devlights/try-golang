/*
net/textproto パッケージのサンプルです。
*/
package main

import (
	"flag"
	"fmt"
	"net"
	"net/textproto"
	"strings"
)

type (
	Args struct {
		IsServer bool
	}
)

const (
	OK = 200
	NG = 400

	CmdGet  = "GET"
	CmdSet  = "SET"
	CmdQuit = "QUIT"
)

var (
	args Args
)

func init() {
	flag.BoolVar(&args.IsServer, "server", false, "server mode")
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var err error

	if args.IsServer {
		err = runServer()
	} else {
		err = runClient()
	}

	if err != nil {
		return err
	}

	return nil
}

func runServer() error {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		return err
	}
	defer l.Close()

	// サンプルなので１回だけ受付
	// クライアントからは
	//   - SET
	//   - GET
	//   - QUIT
	// の３コマンドのシーケンスが来るとする
	conn, err := l.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	// net/textprotoの接続として処理
	// 
	// textproto.NewConn() は、引数に io.ReadWriteCloser を要求しているが
	// net.Conn は、os.Fileと同様に io.ReadWriteCloser を実装している。
	tpConn := textproto.NewConn(conn)
	defer tpConn.Close()

	// ウェルカムメッセージ
	err = tpConn.PrintfLine("%d %s", OK, "WELCOME AVAILABLE COMMANDS: {SET, GET, QUIT}")
	if err != nil {
		return err
	}

	// コマンド処理
	data := make(map[string]string)
	for {
		line, err := tpConn.ReadLine()
		if err != nil {
			return err
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			tpConn.PrintfLine("%d %s", NG, "コマンドが読み取れません")
			continue
		}

		command := strings.ToUpper(parts[0])
		cmdArgs := parts[1:]

		switch command {
		case CmdGet:
			if v, ok := data[cmdArgs[0]]; ok {
				tpConn.PrintfLine("%d %s", OK, v)
			} else {
				tpConn.PrintfLine("%d %s", NG, "KEY NOT FOUND")
			}
		case CmdSet:
			data[cmdArgs[0]] = cmdArgs[1]
			tpConn.PrintfLine("%d %s", OK, "VALUE SET SUCCESSFULLY")
		case CmdQuit:
			tpConn.PrintfLine("%d %s", OK, "BYE")
			return nil
		default:
			tpConn.PrintfLine("%d %s", NG, "UNKNOWN COMMAND")
		}
	}
}

func runClient() error {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		return err
	}
	defer conn.Close()

	tpConn := textproto.NewConn(conn)
	defer tpConn.Close()

	// helper funcs
	var (
		send = func(tp *textproto.Conn, msg string) error {
			fmt.Printf("< %s\n", msg)
			return tp.PrintfLine("%s", msg)
		}
		recv = func(tp *textproto.Conn) error {
			code, msg, err := tp.ReadCodeLine(OK)
			if err != nil {
				return err
			}
			fmt.Printf("%d %s\n", code, msg)

			return nil
		}
	)

	// WELCOME
	{
		err = recv(tpConn)
		if err != nil {
			return err
		}
	}

	// SET
	m := fmt.Sprintf("%s %s %s", CmdSet, "Hello", "World")
	{
		err = send(tpConn, m)
		if err != nil {
			return err
		}

		err = recv(tpConn)
		if err != nil {
			return err
		}
	}

	// GET
	m = fmt.Sprintf("%s %s", CmdGet, "Hello")
	{
		err = send(tpConn, m)
		if err != nil {
			return err
		}

		err = recv(tpConn)
		if err != nil {
			return err
		}
	}

	// 存在しないコマンド
	m = fmt.Sprintf("%s %s", "GOLANG", "HELLO")
	{
		err = send(tpConn, m)
		if err != nil {
			return err
		}

		err = recv(tpConn)
		if err != nil {
			fmt.Printf("%[1]s (%[1]T)\n", err)
		}
	}

	// QUIT
	m = CmdQuit
	{
		err = send(tpConn, m)
		if err != nil {
			return err
		}

		err = recv(tpConn)
		if err != nil {
			return err
		}
	}

	return nil
}
