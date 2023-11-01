package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/devlights/try-golang/examples/singleapp/designpatterns/singleton/defines"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	type (
		Singleton interface {
			GetDef1() string
			GetDef2() string
		}
	)

	var (
		c  = runtime.NumCPU()
		l  = make([]Singleton, 0, c)
		wg sync.WaitGroup
	)

	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			defer wg.Done()
			println(">>> call: GetInstance")
			l = append(l, defines.GetInstance())
		}()
	}

	wg.Wait()
	for _, s := range l {
		fmt.Printf("%p (%s:%s)\n", &s, s.GetDef1(), s.GetDef2())
	}

	return nil
}
