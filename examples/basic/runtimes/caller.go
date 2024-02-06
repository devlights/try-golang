package runtimes

import (
	"runtime"
	"strings"

	"github.com/devlights/gomy/output"
)

// Caller は、 runtime.Caller() のサンプルです.
//
// runtime.Caller() は、呼び出し元のgoroutineのスタックにある関数呼び出しに関するファイルおよび行番号情報を報告してくれる.
//
// # REFERENCES
//   - https://pkg.go.dev/runtime@go1.19.3#Caller
func Caller() error {
	var (
		programCounter uintptr
		file           string
		line           int
		ok             bool
		sentinel       = "try-golang"
	)

	for skip := 0; ; skip++ {
		// skip = 0 は、runtime.Caller() の場合は、この関数を呼び出した部分になる.
		// runtime.Callers() の場合は、runtime.Callers() 自体を表す.
		// 0 の意味が少し違うことに注意.
		programCounter, file, line, ok = runtime.Caller(skip)
		if !ok {
			break
		}

		if !strings.Contains(file, sentinel) {
			break
		}

		output.Stdoutf("[runtime.Caller]", "%v\t%v\t%v\n", programCounter, file, line)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: runtime_caller

	   [Name] "runtime_caller"
	   [runtime.Caller]     8194836    /workspace/try-golang/examples/basic/runtimes/caller.go 29
	   [runtime.Caller]     9119412    /workspace/try-golang/runner/exec.go    52
	   [runtime.Caller]     9121316    /workspace/try-golang/runner/loop.go    126
	   [runtime.Caller]     9120584    /workspace/try-golang/runner/loop.go    86
	   [runtime.Caller]     9124390    /workspace/try-golang/cmd/root.go       66
	   [runtime.Caller]     9125774    /workspace/try-golang/main.go   6

	   [Elapsed] 90.889µs
	*/
}
