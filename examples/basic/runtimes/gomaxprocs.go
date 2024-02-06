package runtimes

import (
	"runtime"

	"github.com/devlights/gomy/output"
)

// GoMaxProcs -- runtime.GOMAXPROCS() のサンプルです。
func GoMaxProcs() error {
	// ------------------------------------------------------------
	// GOMAXPROCS は、Goにて同時実行させるCPUの数を表す
	// Go1.5より前は、デフォルトで 1 だったが、Go1.5よりデフォルトが
	// 動作する環境のCPUの数になっている。
	//
	// https://golang.org/runtime/#GOMAXPROCS
	// https://stackoverflow.com/questions/17853831/what-is-the-gomaxprocs-default-value
	//
	// 動作する環境にてCPUがいくつあるかを取得するには
	//     runtime.NumCPU()
	// を利用する。
	//
	// runtime.GOMAXPROCS() は、設定するよう関数であるが
	// 引数に 0 を指定することにより、現在の値を取得することが出来る。
	// https://golang.org/src/runtime/debug.go?s=533:559#L7
	// ------------------------------------------------------------
	output.Stdoutl("NumCPU", runtime.NumCPU())
	output.Stdoutl("GOMAXPROCS", runtime.GOMAXPROCS(0))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: runtime_gomaxprocs

	   [Name] "runtime_gomaxprocs"
	   NumCPU               16
	   GOMAXPROCS           16


	   [Elapsed] 12.65µs
	*/

}
