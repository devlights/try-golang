package runtimes

import (
	"fmt"
	"runtime"
)

// RuntimeVersion は、runtime.Version() のサンプルです。
func RuntimeVersion() error {
	// runtime.Version() で 現在利用している Go のバージョンが取得できる
	ver := runtime.Version()
	fmt.Println(ver)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: runtime_version

	   [Name] "runtime_version"
	   go1.21.6


	   [Elapsed] 3.21µs
	*/

}
