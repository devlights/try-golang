package stacktrace

import (
	"runtime"

	"github.com/devlights/gomy/output"
)

// RuntimeStack -- runtime.Stack() についてのサンプルです。
//
// REFERENCES
//   - https://pkg.go.dev/runtime#Stack
//   - https://stackoverflow.com/questions/19094099/how-to-dump-goroutine-stacktraces
func RuntimeStack() error {
	// channels
	var (
		chSingle = make(chan []byte, 1)
		chAll    = make(chan []byte, 1)
	)

	// funcs
	var (
		getStack = func(all bool) []byte {
			// From src/runtime/debug/stack.go
			var (
				buf = make([]byte, 1024)
			)

			for {
				n := runtime.Stack(buf, all)
				if n < len(buf) {
					return buf[:n]
				}
				buf = make([]byte, 2*len(buf))
			}
		}
	)

	// runtime.Stack() の 第２引数 に false を渡すと、現在のgoroutineのみが対象
	go func(ch chan<- []byte) {
		defer close(ch)
		ch <- getStack(false)
	}(chSingle)

	// runtime.Stack() の 第２引数 に true  を渡すと、全てのgoroutineのみが対象
	go func(ch chan<- []byte) {
		defer close(ch)
		ch <- getStack(true)
	}(chAll)

	// runtime.Stack(false) の 結果
	for v := range chSingle {
		output.Stdoutl("=== stack-single", string(v))
	}

	// runtime.Stack(true) の 結果
	for v := range chAll {
		output.Stdoutl("=== stack-all   ", string(v))
	}

	return nil
}
