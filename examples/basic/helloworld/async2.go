package helloworld

import (
	"fmt"
	"log"
	"math/rand"
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

type _closer struct {
	wg *sync.WaitGroup
	ch chan<- string
}

func newCloser(wg *sync.WaitGroup, ch chan<- string) *_closer {
	c := new(_closer)

	c.wg = wg
	c.ch = ch

	return c
}

func (me *_closer) run() {
	defer close(me.ch)
	me.wg.Wait()
}

type _printer struct {
	ch <-chan string
}

func newPrinter(ch <-chan string) *_printer {
	p := new(_printer)
	p.ch = ch
	return p
}

func (me *_printer) run() {
	for v := range me.ch {
		log.Println(v)
	}
}

// Async2 -- HelloWorld 非同期版 (2)
func Async2() error {
	log.SetFlags(0)

	var (
		wg    = new(sync.WaitGroup)
		ch    = make(chan string)
		delay = func() time.Duration {
			return time.Duration(rand.Intn(200)) * time.Millisecond
		}
	)

	var (
		hello   = newRunner("hello", wg, ch, delay)
		world   = newRunner("world", wg, ch, delay)
		closer  = newCloser(wg, ch)
		printer = newPrinter(ch)
	)

	wg.Add(2)

	go hello.run()
	go world.run()
	go closer.run()

	printer.run()

	return nil
}
