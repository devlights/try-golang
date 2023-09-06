package async3

import (
	"fmt"
	"time"
)

type _runner struct {
	name  string
	ch    chan string
	count int
	delay func() time.Duration
}

func newRunner(name string, count int, delay func() time.Duration) *_runner {
	r := new(_runner)

	r.name = name
	r.ch = make(chan string)
	r.count = count
	r.delay = delay

	return r
}

func (me *_runner) run() {
	go func() {
		defer close(me.ch)

		for i := 0; i < me.count; i++ {
			d := me.delay()
			time.Sleep(d)
			me.ch <- fmt.Sprintf("%d:%s (%v)", i, me, d)
		}
	}()
}

func (me *_runner) String() string {
	return me.name
}
