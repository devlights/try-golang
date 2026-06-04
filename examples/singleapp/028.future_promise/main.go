package main

import (
	"log"
	"sync"
	"time"
)

type (
	// Future[T] は、「今はまだ得られていないが将来得られるはずの入力」を表します。
	Future[T any] struct {
		value T
		wait  chan struct{}
	}

	// Promise[T] は、「将来値を提供するという約束」を表します。
	Promise[T any] struct {
		f *Future[T]
	}
)

func (me *Future[T]) IsDone() bool {
	select {
	case <-me.wait:
		return true
	default:
		return false
	}
}

func (me *Future[T]) set(v T) {
	if me.IsDone() {
		return
	}

	me.value = v
	close(me.wait)
}

func (me *Future[T]) Get() T {
	<-me.wait
	return me.value
}

func (me *Promise[T]) Submit(v T) {
	me.f.set(v)
}

func NewPromise[T any]() (*Future[T], *Promise[T]) {
	f := &Future[T]{wait: make(chan struct{})}
	p := &Promise[T]{f: f}

	return f, p
}

func init() {
	log.SetFlags(log.Ltime)
}

func main() {
	log.Println("START")
	defer log.Println("END  ")

	var (
		f1, p1 = NewPromise[int]()
		f2, p2 = NewPromise[string]()
		wg     sync.WaitGroup
	)

	wg.Add(3)

	// f1に依存する処理
	//
	// f2の状態に関係なく、f1が完了したら完了する
	go func(f *Future[int]) {
		defer wg.Done()
		log.Printf("f1=%v", f1.Get())
	}(f1)

	// f2に依存する処理
	//
	// f1の状態に関係なく、f2が完了したら完了する
	go func(f *Future[string]) {
		defer wg.Done()
		log.Printf("f2=%v", f2.Get())
	}(f2)

	// f1とf2に依存する処理
	//
	// f1, f2の２つのFutureが完了しないと処理が完了しない
	go func(f1 *Future[int], f2 *Future[string]) {
		defer wg.Done()
		log.Printf("f1=%v\tf2=%v", f1.Get(), f2.Get())
	}(f1, f2)

	// 1秒後にf1に値を提供する
	//
	// このPromiseが値を提供するまで対応するFuture(f1)は結果を返さない
	go func(p *Promise[int]) {
		time.Sleep(1 * time.Second)
		p.Submit(999)
	}(p1)

	// 3秒後にf2に値を提供する
	//
	// このPromiseが値を提供するまで対応するFuture(f2)は結果を返さない
	go func(p *Promise[string]) {
		time.Sleep(3 * time.Second)
		p.Submit("hello world")
	}(p2)

	wg.Wait()

	/*
	   $ task -d examples/singleapp/future_promise/
	   task: [default] go run main.go
	   08:19:13 START
	   08:19:14 f1=999
	   08:19:16 f2=hello world
	   08:19:16 f1=999 f2=hello world
	   08:19:16 END
	*/

}
