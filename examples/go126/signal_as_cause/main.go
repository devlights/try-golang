package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type (
	_Args struct {
		sigterm bool
	}
)

var (
	args  _Args
	block = make(chan bool)
)

func main() {
	log.SetFlags(0)

	flag.BoolVar(&args.sigterm, "term", false, "send SIGTERM")
	flag.Parse()

	if err := run(context.Background()); err != nil {
		panic(err)
	}
}

func run(pCtx context.Context) error {
	//
	// Go 1.26 にて、signal.NotifyContext でシグナルを受けた際に
	// context.Cause(ctx) とすると、どのシグナルを受けたかが文字列であるが分かるようになった。
	// （以前は context canceled だった。Go 1.26 では xxx signal received となる。）
	//
	// エラーとしては signal.signalError となっており、実質文字列でしか利用できないが有用。
	// (https://cs.opensource.google/go/go/+/refs/tags/go1.26.0:src/os/signal/signal.go;l=340)
	//
	// > type signalError string
	//
	// ちなみに、ctx.Err() と context.Cause(ctx) の違いは以下のようになる。
	//
	// - ctx.Err()          は「何が起きたか？」を表す
	// - context.Cause(ctx) は「何故起きたか？」を表す
	//
	// ctx.Err() は、２種類のエラーしか返さない
	//   - context.Canceled
	//   - context.DeadlineExceeded
	//
	// 「何かが起きたのか」は分かるが「何故キャンセルされた？」「何故タイムアウトした？」が分からない。
	// その点を補うのが context.Cause(ctx) の方。こちらは任意のエラーを設定出来る。
	//
	// context.Cause()に設定するためには、ctx 作成時に context.WithXXXXCause() を利用する必要がある。
	// ここで取得できる context.CancelCauseFunc を使って、任意のエラーを設定する。
	//

	var (
		ctx  context.Context
		stop context.CancelFunc
	)
	ctx, stop = signal.NotifyContext(pCtx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 自プロセス取得
	var (
		self *os.Process
		err  error
	)
	if self, err = os.FindProcess(os.Getpid()); err != nil {
		return err
	}

	// 自分にシグナルを送る
	var (
		sig = syscall.SIGINT
	)
	if args.sigterm {
		sig = syscall.SIGTERM
	}
	if err = self.Signal(sig); err != nil {
		return err
	}

	// ハンドリング
	select {
	case <-block:
	case <-ctx.Done():
		log.Printf("Err=(%[1]v,%[1]T), Cause=(%[2]v,%[2]T)", ctx.Err(), context.Cause(ctx))
	}

	return nil
}
