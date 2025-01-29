package runtimes

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// RuntimeVersion は、runtime.Version() のサンプルです。
func RuntimeVersion() error {
	// runtime.Version() で 現在利用している Go のバージョンが取得できる
	ver := runtime.Version()
	fmt.Println(ver)

	// debug.ReadBuildInfo() からも取得することが出来る
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return fmt.Errorf("failed: debug.ReadBuildInfo()")
	}
	fmt.Println(info.GoVersion)

	return nil

	/*
		$ task
		task: [build] go build -o "/workspace/try-golang/try-golang" .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: runtime_version

		[Name] "runtime_version"
		go1.23.5
		go1.23.5


		[Elapsed] 48.21µs
	*/

}
