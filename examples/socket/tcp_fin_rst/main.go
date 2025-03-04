package main

import (
	"flag"
	"net"
)

type (
	Args struct {
		IsServer bool
		UseRst   bool
	}
)

var (
	args Args
)

func init() {
	flag.BoolVar(&args.IsServer, "server", false, "サーバモードで起動")
	flag.BoolVar(&args.UseRst, "rst", false, "RST送信による強制切断を使用")
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var (
		err error
	)
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
	var (
		l   net.Listener
		err error
	)
	l, err = net.Listen("tcp", ":8888")
	if err != nil {
		return err
	}
	defer l.Close()

	// サンプルなので１回のみ接続を受付
	var (
		conn net.Conn
	)
	conn, err = l.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	// 適当にデータを送信しクライアントが切ってきたら終わり
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		return err
	}

	var (
		buf = make([]byte, 10)
		n   int
	)
	for {
		clear(buf)

		n, err = conn.Read(buf)
		if n == 0 || err != nil {
			break
		}
	}

	if args.UseRst {
		// RST送信するために SO_LINGER を設定
		// Goの場合 *net.TCPConn に SetLinger メソッドが用意されている。
		var (
			tcpConn   *net.TCPConn
			ok        bool
			lingerSec = 0
		)
		tcpConn, ok = conn.(*net.TCPConn)
		if ok {
			// $ go doc net.tcpconn.setlinger
			//
			// > SetLinger sets the behavior of Close on a connection which still has data waiting to be sent or to be acknowledged.
			// > If sec < 0 (the default), the operating system finishes sending the data in the background.
			// > If sec == 0, the operating system discards any unsent or unacknowledged data.
			// > If sec > 0, the data is sent in the background as with sec < 0. On some operating systems including Linux,
			// > this may cause Close to block until all data has been sent or discarded.
			// > On some operating systems after sec seconds have elapsed any remaining unsent data may be discarded.
			tcpConn.SetLinger(lingerSec)
		}
	}

	return nil
}

func runClient() error {
	var (
		conn net.Conn
		err  error
	)
	conn, err = net.Dial("tcp", "localhost:8888")
	if err != nil {
		return err
	}
	defer func() {
		conn.Close()
	}()

	// データを受信したら切断
	var (
		buf = make([]byte, 10)
	)
	_, err = conn.Read(buf)
	if err != nil {
		return err
	}

	if args.UseRst {
		tcpConn, ok := conn.(*net.TCPConn)
		if ok {
			tcpConn.SetLinger(0)
		}
	}

	return nil
}
