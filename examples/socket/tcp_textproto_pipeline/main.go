package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"net/textproto"
	"strconv"
	"strings"
	"time"
)

type (
	Args struct {
		IsServer bool
	}
)

const (
	OK = 200
	NG = 400

	CmdAdd   = "ADD"
	CmdTotal = "TOTAL"
	CmdQuit  = "QUIT"
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

	switch {
	case args.IsServer:
		err = runServer()
	default:
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
	//   - ADD, ADD, ADD... (textprotoのPipelineを使って複数回分をバッチ送信)
	//   - VAL
	//   - QUIT
	// のシーケンスが来るとする
	conn, err := l.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	tpConn := textproto.NewConn(conn)
	defer tpConn.Close()

	var total int64
	for {
		line, err := tpConn.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return err
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			tpConn.PrintfLine("%d %s", NG, "COMMAND COULD NOT BE READ")
			continue
		}

		command := strings.ToUpper(parts[0])
		cmdArgs := parts[1:]
		switch command {
		case CmdAdd:
			v, err := strconv.Atoi(cmdArgs[0])
			if err != nil {
				tpConn.PrintfLine("%d %s", NG, "VALUE SHOULD BE A NUMBER")
				continue
			}

			total += int64(v)
			tpConn.PrintfLine("%d %s %d", OK, "ADDED", v)
		case CmdTotal:
			tpConn.PrintfLine("%d %s", OK, strconv.FormatInt(total, 10))
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

	//
	// パイプラインを使って複数のリクエストを順序保証しながらバッチ処理
	//
	const numRequests = 5
	ids := make([]uint, 0, numRequests)
	for range numRequests {
		// IDを採番
		id := tpConn.Next()

		// 取得したIDを使ってリクエスト
		tpConn.StartRequest(id)
		{
			msg := fmt.Sprintf("%s %d", CmdAdd, time.Now().Nanosecond())
			err = send(conn, tpConn, msg)
			if err != nil {
				return err
			}
		}
		tpConn.EndRequest(id)

		ids = append(ids, id)
	}

	for _, id := range ids {

		// リクエスト時に採番されたIDを使って応答受信
		tpConn.StartResponse(id)
		{
			err = recv(conn, tpConn)
			if err != nil {
				return err
			}
		}
		tpConn.EndResponse(id)
	}

	// 合計を聞く
	err = send(conn, tpConn, CmdTotal)
	if err != nil {
		return err
	}

	err = recv(conn, tpConn)
	if err != nil {
		return err
	}

	// 終わり
	err = send(conn, tpConn, CmdQuit)
	if err != nil {
		return err
	}

	err = recv(conn, tpConn)
	if err != nil {
		return err
	}

	return nil
}

func send(conn net.Conn, tpConn *textproto.Conn, msg string) error {
	err := conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return err
	}

	fmt.Printf("< %s\n", msg)
	err = tpConn.PrintfLine("%s", msg)
	if err != nil {
		return err
	}

	return nil
}

func recv(conn net.Conn, tpConn *textproto.Conn) error {
	err := conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		return err
	}

	code, msg, err := tpConn.ReadCodeLine(OK)
	if err != nil {
		return err
	}
	fmt.Printf("%d %s\n", code, msg)

	return nil
}
