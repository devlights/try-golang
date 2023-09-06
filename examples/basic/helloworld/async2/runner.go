package async2

import (
	"fmt"
	"sync"
	"time"
)

const (
	LoopCount = 20
)

type _runner struct {
	name  string
	wg    *sync.WaitGroup
	ch    chan<- string
	delay func() time.Duration
}

func newRunner(name string, wg *sync.WaitGroup, ch chan<- string, delay func() time.Duration) *_runner {
	r := new(_runner)

	r.name = name
	r.wg = wg
	r.ch = ch
	r.delay = delay

	return r
}

func (me *_runner) String() string {
	return me.name
}

func (me *_runner) run() {
	defer me.wg.Done()
	for i := 0; i < LoopCount; i++ {
		d := me.delay()
		time.Sleep(d)
		me.ch <- fmt.Sprintf("%d:%s (%v)", i, me, d)
	}
}
