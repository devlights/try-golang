package main

import (
	"flag"
	"fmt"
	"net"
	"net/rpc"
)

type (
	Args struct {
		IsServer bool
		X        int64
		Y        int64
	}
	RpcArgs struct {
		X int64
		Y int64
	}
	RpcReply struct {
		Result int64
	}
	Service struct{}
)

func (me *Service) Multiply(req *RpcArgs, res *RpcReply) error {
	if req == nil {
		return fmt.Errorf("req is nil")
	}

	result := req.X * req.Y
	res.Result = result

	return nil
}

var (
	args Args
)

func init() {
	flag.BoolVar(&args.IsServer, "server", false, "server mode")
	flag.Int64Var(&args.X, "x", 0, "X")
	flag.Int64Var(&args.Y, "y", 0, "Y")
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var err error
	switch args.IsServer {
	case true:
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
	//
	// RPCとして公開するサービス登録
	//
	service := new(Service)
	rpc.Register(service)

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		return err
	}
	defer l.Close()

	rpc.Accept(l)
	// rpc.Accept(net.Listener) は以下と同じ。
	//
	// for {
	// 	conn, err := l.Accept()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	go rpc.ServeConn(conn)
	// }

	return nil
}

func runClient() error {
	client, err := rpc.Dial("tcp", ":8888")
	if err != nil {
		return err
	}
	defer client.Close()

	//
	// RPC 呼び出し
	//   サービスメソッドの引数は arg, reply ともにポインタで渡すこと
	//
	const (
		serviceMethod = "Service.Multiply"
	)
	var (
		a = RpcArgs{X: args.X, Y: args.Y}
		r RpcReply
	)
	err = client.Call(serviceMethod, &a, &r)
	if err != nil {
		return err
	}

	fmt.Printf("%d * %d = %d\n", args.X, args.Y, r.Result)

	return nil
}
