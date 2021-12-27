package bitop

import (
	"context"
	"time"

	"github.com/devlights/gomy/output"
)

const (
	DoneProducer int = 1 << iota
	DoneCompleter
	DoneConsumer1
	DoneConsumer2
	DoneAll = DoneProducer | DoneCompleter | DoneConsumer1 | DoneConsumer2
	None    = 0
)

type (
	producer  chan<- interface{}
	consumer  <-chan interface{}
	completer chan<- interface{}
)

func (me producer) put(n int) context.Context {
	ctx, cxl := context.WithCancel(context.Background())
	go func() {
		defer cxl()
		for i := 0; i < n; i++ {
			me <- i
		}
	}()
	return ctx
}

func (me consumer) take(prefix string) context.Context {
	ctx, cxl := context.WithCancel(context.Background())
	go func() {
		defer cxl()
		for v := range me {
			output.Stderrf(prefix, "%v\n", v)
		}
	}()
	return ctx
}

func (me completer) completeWhen(doneCtx context.Context) context.Context {
	ctx, cxl := context.WithTimeout(doneCtx, 1*time.Minute)
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

// BitFlags -- ビットフラグのサンプルです.
func BitFlags() error {
	// initialize jobs
	var (
		ch   = make(chan interface{})
		p    = producer(ch)
		c1   = consumer(ch)
		c2   = consumer(ch)
		comp = completer(ch)
	)

	// start tasks
	var (
		ctxP    = p.put(30)
		ctxComp = comp.completeWhen(ctxP)
		ctx1C   = c1.take("c-1")
		ctx2C   = c2.take("c-2")
		status  = None
	)

	// wait until all jobs is done
LOOP:
	for {
		select {
		case <-ctxP.Done():
			status |= DoneProducer
		case <-ctxComp.Done():
			status |= DoneCompleter
		case <-ctx1C.Done():
			status |= DoneConsumer1
		case <-ctx2C.Done():
			status |= DoneConsumer2
		}

		if status == DoneAll {
			break LOOP
		}
	}

	return nil
}
