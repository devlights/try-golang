/*
GOTRACEBACKの値をプログラムから設定するサンプル

runtime/debug.SetTraceback("all") の様に設定できる。

# REFERENCES
  - https://pkg.go.dev/runtime/debug#SetTraceback
  - https://pkg.go.dev/runtime#hdr-Environment_Variables
  - https://go.dev/doc/godebug
*/
package main

import (
	"flag"
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

type (
	Args struct {
		Traceback string
	}
)

var (
	args Args
)

func init() {
	flag.StringVar(&args.Traceback, "traceback", "single", "GOTRACEBACK (none, single, all, system, crach, wer)")
}

func main() {
	flag.Parse()

	if args.Traceback == "" {
		args.Traceback = "single"
	}
	debug.SetTraceback(args.Traceback)

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	//
	// 2秒経過後に意図的にパニックさせる
	//
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(3)

	go func(ch chan<- int) {
		defer wg.Done()
		for i := range 100 {
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
	}(ch)
	go func(ch <-chan int) {
		defer wg.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}(ch)
	go func(ch chan<- int) {
		defer wg.Done()
		defer close(ch)
		time.Sleep(2 * time.Second)
	}(ch)

	wg.Wait()

	return nil
}
