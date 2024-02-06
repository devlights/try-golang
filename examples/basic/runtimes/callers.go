package runtimes

import (
	"runtime"

	"github.com/devlights/gomy/output"
)

// Callers は、 runtime.Callers() のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/runtime@go1.19.3#Callers
func Callers() error {
	var (
		skip       = 0
		pc         = make([]uintptr, 10)
		frameCount = runtime.Callers(skip, pc)
		frames     = runtime.CallersFrames(pc[:frameCount])
	)

	output.Stdoutf("[runtime.CallersFrames]", "frames=%v\n", frameCount)

	// *runtime.Frames を イテレーション して値を取得
	for {
		frame, more := frames.Next()

		output.Stdoutf("[runtime.Frames]", "%v\t%v\t%v\n", frame.PC, frame.File, frame.Line)

		if !more {
			break
		}
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: runtime_callers

	   [Name] "runtime_callers"
	   [runtime.CallersFrames] frames=9
	   [runtime.Frames]     8195016    /home/gitpod/go/src/runtime/extern.go   308
	   [runtime.Frames]     8194998    /workspace/try-golang/examples/basic/runtimes/callers.go        17
	   [runtime.Frames]     9119412    /workspace/try-golang/runner/exec.go    52
	   [runtime.Frames]     9121316    /workspace/try-golang/runner/loop.go    126
	   [runtime.Frames]     9120348    /workspace/try-golang/runner/loop.go    79
	   [runtime.Frames]     9124390    /workspace/try-golang/cmd/root.go       66
	   [runtime.Frames]     9125774    /workspace/try-golang/main.go   6
	   [runtime.Frames]     4452858    /home/gitpod/go/src/runtime/proc.go     267
	   [runtime.Frames]     4654400    /home/gitpod/go/src/runtime/asm_amd64.s 1650

	   [Elapsed] 140.35µs
	*/
}
