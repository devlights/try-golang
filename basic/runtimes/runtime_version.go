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
}
