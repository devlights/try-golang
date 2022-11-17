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
}
