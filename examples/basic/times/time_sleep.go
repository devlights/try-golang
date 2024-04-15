package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// Sleep は、time.Sleep() のサンプルです。
//
// > Sleep pauses the current goroutine for at least the duration d. A negative or zero duration causes Sleep to return immediately.
//
// > スリープは、少なくとも継続時間dの間、現在のゴルーチンを一時停止します。継続時間が負またはゼロの場合、スリープは即座に戻ります。
//
// # REFERENCES
//   - https://pkg.go.dev/time@go1.22.2#Sleep
func Sleep() error {
	const (
		timeFormat = time.TimeOnly + ".000"
	)

	//
	// time.Sleep() は、他の言語のSleep関数と同じ挙動。
	// 指定した時間分、現在のスレッドをブロックする。
	//
	output.Stdoutl("[begin]", time.Now().Format(timeFormat))
	defer func() { output.Stdoutl("[end]", time.Now().Format(timeFormat)) }()

	time.Sleep(3 * time.Second)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_sleep

	   [Name] "time_sleep"
	   [begin]              07:32:51.959
	   [end]                07:32:54.960


	   [Elapsed] 3.000651935s
	*/

}
