package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"
)

type (
	// Event はイベント駆動処理の単一イベントを表す。
	Event struct {
		id  int64
		val string
	}
)

// String は Event の文字列表現を返す。
func (me Event) String() string {
	return fmt.Sprintf("id=%d,val=%q", me.id, me.val)
}

type (
	// Proc はイベントを生成し続けるプロセッサ。
	Proc struct {
		done      chan struct{}  // done は Close() 呼び出し時に close される終了通知チャネル。
		events    chan Event     // events は生成したイベントを購読者に届けるチャネル。
		wg        sync.WaitGroup // wg は内部ゴルーチンの完了を追跡する。
		closeOnce sync.Once      // closeOnce は Close 時の処理を一度のみ実行させるためのもの。
	}
)

// NewProc は Proc を生成して返す。
func NewProc() *Proc {
	var (
		done = make(chan struct{})
		ch   = make(chan Event, 0xFF)
	)
	return &Proc{
		done:   done,
		events: ch,
	}
}

// Run はイベント生成ゴルーチンを起動する。
func (me *Proc) Run(pCtx context.Context) {
	me.wg.Go(func() {
		defer close(me.events)

		me.events <- Event{id: 0, val: "first"}

		var (
			t = time.NewTicker(1 * time.Second)
		)
		defer t.Stop()

		for i := 1; ; i++ {
			select {
			case <-me.done:
				me.events <- Event{id: -1, val: "last"}
				return
			case <-pCtx.Done():
				me.events <- Event{id: -9, val: "last"}
				return
			case tick := <-t.C:
				e := Event{id: tick.Unix(), val: fmt.Sprintf("hello-%d", i)}
				me.events <- e
			}
		}
	})
}

// Events は読み取り専用のイベントチャネルを返す。
func (me *Proc) Events() <-chan Event {
	return me.events
}

// Close は Proc を停止する。
func (me *Proc) Close() error {
	me.closeOnce.Do(func() {
		close(me.done)
		me.wg.Wait()
	})

	return nil
}

type (
	// Args は、プログラム引数を表す。
	Args struct {
		// timeout は、実行時間を指定できるオプション。(デフォルト値あり)
		timeout time.Duration
	}
)

const (
	defaultTimeout = 5 * time.Second
)

var (
	args Args
)

func init() {
	log.SetFlags(log.Lmicroseconds)
	flag.DurationVar(&args.timeout, "t", defaultTimeout, "timeout")
}

func main() {
	flag.Parse()

	var (
		ctx = context.Background()
		err error
	)
	if err = run(ctx); err != nil {
		log.Panic(err)
	}
}

func run(pCtx context.Context) error {
	var (
		ctx, cxl = context.WithTimeout(pCtx, args.timeout)
		proc     = NewProc()
	)
	defer cxl()

	proc.Run(ctx)

	log.Println("START")
	defer func() { log.Println("DONE") }()

	go func() {
		// サンプルなので少し待機させたらクローズさせる
		// 実際はシグナルなどを受けてクローズさせたりするのが一般的。
		time.Sleep(defaultTimeout)
		proc.Close()
	}()

	for v := range proc.Events() {
		log.Println(v)
	}

	return nil
}
