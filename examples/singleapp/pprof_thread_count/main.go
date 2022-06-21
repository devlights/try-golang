/*
	Goが内部で利用しているスレッド数を表示するサンプルです。

	結果は例えば以下のようになります。(結果は環境によって異なります。)

		$ go install github.com/go-task/task/v3/cmd/task@latest
		$ go install honnef.co/go/tools/cmd/staticcheck@latest
		$ task
		task: [run] go fmt
		task: [run] go vet ./...
		task: [run] staticcheck ./...
		task: [run] go run main.go
		[BEFORE] lock-thread=false      thread count=5
		[AFTER ] lock-thread=false      thread count=5
		[BEFORE] lock-thread=true       thread count=5
		[AFTER ] lock-thread=true       thread count=15

	REFERENCES:
		- https://blog.rahuldev.in/how-to-implement-concurrency-and-parallelism-in-go
*/
package main

import (
	"fmt"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

func fn(lockThread bool, wg *sync.WaitGroup) {
	defer wg.Done()

	if lockThread {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	}

	time.Sleep(1 * time.Second)
}

func main() {
	const (
		NUM_GOROUTINE = 10
	)

	var (
		threadProfile = pprof.Lookup("threadcreate")
	)

	for _, lockThread := range []bool{false, true} {
		var (
			wg sync.WaitGroup
		)

		fmt.Printf("[BEFORE] lock-thread=%v\tthread count=%d\n", lockThread, threadProfile.Count())

		wg.Add(NUM_GOROUTINE)
		for i := 0; i < NUM_GOROUTINE; i++ {
			go fn(lockThread, &wg)
		}
		wg.Wait()

		fmt.Printf("[AFTER ] lock-thread=%v\tthread count=%d\n", lockThread, threadProfile.Count())
	}
}
