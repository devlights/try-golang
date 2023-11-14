package bitop

import (
	"context"
	"log"
	"math/rand"
	"time"
)

// ---------------------------------------
// Enums
// ---------------------------------------

type doneStatus uint8

const (
	DoneProducer doneStatus = 1 << iota
	DoneCompleter
	DoneConsumer1
	DoneConsumer2
	DoneAll = DoneProducer | DoneCompleter | DoneConsumer1 | DoneConsumer2
	None    = 0
)

// ---------------------------------------
// Vars
// ---------------------------------------
var (
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// ---------------------------------------
// Types
// ---------------------------------------

type (
	producer  chan<- interface{}
	consumer  <-chan interface{}
	completer chan<- interface{}
	notifier  chan<- doneStatus
	reporter  <-chan doneStatus
)

func (me producer) put(n int) context.Context {
	ctx, cxl := context.WithCancel(context.Background())
	go func() {
		defer cxl()
		for i := 0; i < n; i++ {
			me <- i
			time.Sleep(time.Duration(rnd.Intn(10)) * time.Millisecond)
		}
	}()
	return ctx
}

func (me consumer) take(prefix string) context.Context {
	ctx, cxl := context.WithCancel(context.Background())
	go func() {
		defer cxl()
		for v := range me {
			log.Printf("%s: %v\n", prefix, v)
			time.Sleep(time.Duration(rnd.Intn(50)) * time.Millisecond)
		}
	}()
	return ctx
}

func (me completer) completeWhen(doneCtx context.Context) context.Context {
	ctx, cxl := context.WithTimeout(doneCtx, 3*time.Second)
	go func() {
		defer cxl()
		select {
		case <-doneCtx.Done():
		case <-ctx.Done():
		}

		close(me)
	}()
	return ctx
}

func (me notifier) notify(status doneStatus) {
	ctx, cxl := context.WithTimeout(context.Background(), 1*time.Second)
	defer cxl()

	select {
	case <-ctx.Done():
	case me <- status:
	}
}

func (me notifier) stop() {
	close(me)
}

func (me reporter) start() context.Context {
	ctx, cxl := context.WithCancel(context.Background())

	fn := func(b bool) string {
		if b {
			return "DONE   "
		} else {
			return "RUNNING"
		}
	}

	go func() {
		defer cxl()
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case v, ok := <-me:
				if !ok {
					break LOOP
				}
				var (
					pDone    = (v & DoneProducer) == DoneProducer
					compDone = (v & DoneCompleter) == DoneCompleter
					c1Done   = (v & DoneConsumer1) == DoneConsumer1
					c2Done   = (v & DoneConsumer2) == DoneConsumer2
				)

				log.Printf(
					"Producer:%v  Completer:%v  Consumer1:%v  Consumer2: %v\n",
					fn(pDone), fn(compDone), fn(c1Done), fn(c2Done))
			}
		}
	}()
	return ctx
}

// BitFlags -- ビットフラグのサンプルです.
func BitFlags() error {
	log.SetFlags(0)

	// initialize jobs
	var (
		ch   = make(chan interface{}, 30)
		p    = producer(ch)
		c1   = consumer(ch)
		c2   = consumer(ch)
		comp = completer(ch)
		nCh  = make(chan doneStatus)
		n    = notifier(nCh)
		r    = reporter(nCh)
	)

	// start tasks
	var (
		ctxP    = p.put(30)
		ctxComp = comp.completeWhen(ctxP)
		ctx1C   = c1.take("c-1")
		ctx2C   = c2.take("c-2")
		ctxDone = waitUntil(n, ctxP.Done(), ctxComp.Done(), ctx1C.Done(), ctx2C.Done())
		ctxR    = r.start()
	)

	// wait until all jobs is done
	<-ctxDone.Done()
	n.stop()
	<-ctxR.Done()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: bitop_bitflags

	   [Name] "bitop_bitflags"
	   c-1: 0
	   c-2: 1
	   c-2: 2
	   c-1: 3
	   c-1: 4
	   c-1: 5
	   c-2: 6
	   c-1: 7
	   c-1: 8
	   c-2: 9
	   c-1: 10
	   c-2: 11
	   c-1: 12
	   c-2: 13
	   c-1: 14
	   c-2: 15
	   c-1: 16
	   Producer:DONE     Completer:RUNNING  Consumer1:RUNNING  Consumer2: RUNNING
	   Producer:DONE     Completer:DONE     Consumer1:RUNNING  Consumer2: RUNNING
	   c-2: 17
	   c-1: 18
	   c-1: 19
	   c-2: 20
	   c-1: 21
	   c-2: 22
	   c-1: 23
	   c-1: 24
	   c-2: 25
	   c-1: 26
	   c-1: 27
	   c-1: 28
	   c-2: 29
	   Producer:DONE     Completer:DONE     Consumer1:DONE     Consumer2: RUNNING
	   Producer:DONE     Completer:DONE     Consumer1:DONE     Consumer2: DONE


	   [Elapsed] 319.053956ms
	*/

}

func waitUntil(n notifier, doneP, doneComp, done1C, done2C <-chan struct{}) context.Context {
	ctx, cxl := context.WithCancel(context.Background())
	go func() {
		defer cxl()

		var status doneStatus
	LOOP:
		for {
			select {
			case <-doneP:
				status |= DoneProducer
				doneP = nil
			case <-doneComp:
				status |= DoneCompleter
				doneComp = nil
			case <-done1C:
				status |= DoneConsumer1
				done1C = nil
			case <-done2C:
				status |= DoneConsumer2
				done2C = nil
			}

			n.notify(status)

			if status == DoneAll {
				break LOOP
			}
		}
	}()
	return ctx
}
