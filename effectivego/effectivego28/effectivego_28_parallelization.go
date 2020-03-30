package effectivego28

import (
	"fmt"
	"sync"
	"time"

	"github.com/devlights/gomy/output"
)

// Interface defines
type (
	chProc interface {
		calc(doneSignalCh chan<- struct{})
	}

	wgProc interface {
		calc(wg *sync.WaitGroup)
	}
)

// Struct defines
type (
	baseprc struct {
		name string
		t1   time.Duration
	}

	chprc struct {
		baseprc
	}

	wgprc struct {
		baseprc
	}
)

func newChProc(name string, t1 time.Duration) chProc {
	return &chprc{
		baseprc{
			name: name,
			t1:   t1,
		},
	}
}

func newWgProc(name string, t1 time.Duration) wgProc {
	return &wgprc{
		baseprc{
			name: name,
			t1:   t1,
		},
	}
}

// calc -- impl chProc::calc
func (c *chprc) calc(ch chan<- struct{}) {
	defer func(ch chan<- struct{}) {
		// 処理完了を通知
		ch <- struct{}{}
	}(ch)

	output.Stdoutf("[処理中]", "%v\n", c.name)
	<-time.After(c.t1)
	output.Stdoutf("[処理完]", "%v\n", c.name)
}

// calc -- impl wgProg::calc
func (w *wgprc) calc(wg *sync.WaitGroup) {
	defer func(wg *sync.WaitGroup) {
		// 処理完了を通知
		wg.Done()
	}(wg)

	output.Stdoutf("[処理中]", "%v\n", w.name)
	<-time.After(w.t1)
	output.Stdoutf("[処理完]", "%v\n", w.name)
}

// Parallelization -- Effective Go - Parallelization の 内容についてのサンプルです。
func Parallelization() error {
	/*
		https://golang.org/doc/effective_go.html#parallel

		- 計算を独立して実行できる部分に分割することができれば並列化することが可能となる (全ての場合ではない)
		  - 後、必要なのは各部分が完了したときに信号を送ること。これはチャネルを利用することで可能である。
		  - それぞれの並行処理に対して sync.WaitGroup を渡しても良い。大事なのは全ての結果を待機することが可能な状態を作ること
	*/
	var (
		t1 = []time.Duration{
			1 * time.Second,
			2 * time.Second,
			3 * time.Second,
		}
	)

	var (
		wg = sync.WaitGroup{}
		ch = make(chan struct{}, len(t1))
	)

	for i := 0; i < len(t1); i++ {
		proc := newChProc(fmt.Sprintf("chproc-%02d", i), t1[i])
		go proc.calc(ch)
	}

	for i := 0; i < len(t1); i++ {
		// 完了するたびにゴルーチン側から値が送信されてくるので受信するようにする
		// 起動したゴルーチンと同じ数の受信が行えたら、全非同期処理が完了したことになる
		<-ch
	}

	output.Stdoutl("[ChProc]", "All Done")
	output.StdoutHr()

	for i := 0; i < len(t1); i++ {
		proc := newWgProc(fmt.Sprintf("wgproc-%02d", i), t1[i])
		wg.Add(1)

		go proc.calc(&wg)
	}

	wg.Wait()

	output.Stdoutl("[WgProc]", "All Done")

	return nil
}
