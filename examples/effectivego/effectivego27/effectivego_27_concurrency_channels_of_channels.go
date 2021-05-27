package effectivego27

import (
	"time"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/output"
)

type (
	calclator interface {
		calc(*request)
	}

	request struct {
		value    int
		resultCh chan int
	}

	client struct {
		calc calclator
	}

	server struct {
		done  chan struct{}
		reqCh chan interface{}
	}
)

func newServer() *server {
	return &server{}
}

func newClient(c calclator) *client {
	return &client{calc: c}
}

func (c *client) send(req *request) <-chan int {
	go func() {
		c.calc.calc(req)
	}()

	return req.resultCh
}

func (c *client) close() {
	c.calc = nil
}

func (s *server) start() <-chan struct{} {
	s.done = make(chan struct{})
	s.reqCh = make(chan interface{})
	go func(done chan struct{}, ch chan interface{}) {
		defer close(ch)

		for v := range chans.OrDone(done, ch) {
			go func(reqValue interface{}) {
				req, ok := reqValue.(*request)
				if !ok {
					return
				}

				defer close(req.resultCh)

				<-time.After(10 * time.Millisecond)
				req.resultCh <- req.value + 1
			}(v)
		}
	}(s.done, s.reqCh)

	return s.done
}

func (s *server) calc(r *request) {
	s.reqCh <- r
}

func (s *server) stop() {
	defer close(s.done)
}

// ChannelsOfChannels -- Effective Go - Channels of channels の 内容についてのサンプルです。
func ChannelsOfChannels() error {
	/*
		https://golang.org/doc/effective_go.html#chan_of_chan

		Go において、チェネルは第一級市民のオブジェクトである。
		なので、関数の引数や構造体のフィールドなど、どこでも利用できるオブジェクトとなっている。
	*/
	// サーバもどきとクライアントもどきを生成
	s := newServer()
	c := newClient(s)

	// サーバもどきを開始
	serverDone := s.start()

	// リクエストをクライアントもどき経由で送り込む
	resultChList := make([]<-chan interface{}, 0, 0)
	for i := 0; i < 20; i++ {
		resultCh := c.send(&request{
			value:    i,
			resultCh: make(chan int),
		})

		resultChList = append(resultChList, chans.FromIntCh(resultCh))
	}

	// 結果を取得
	fanInCh := chans.FanIn(serverDone, resultChList...)
	for v := range fanInCh {
		output.Stderrl("[result]", v)
	}

	// 停止
	c.close()
	s.stop()
	<-serverDone

	return nil
}
