package runtimes

import (
	"runtime"

	"github.com/devlights/gomy/output"
)

// NumCpu -- runtime.NumCPU() のサンプルです。
//
// # REFERENCES
//   - https://dev.to/freakynit/the-very-useful-runtime-package-in-golang-5b16
func NumCpu() error {
	//
	// runtime.NumCPU() は、プログラムが動作しているマシンのCPUのコア数を返してくれる
	//

	var (
		cpus = runtime.NumCPU()
	)
	output.Stdoutl("[Number of CPU cores]", cpus)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: runtime_numcpu

	   [Name] "runtime_numcpu"
	   [Number of CPU cores] 16

	   [Elapsed] 8.79µs
	*/
}
