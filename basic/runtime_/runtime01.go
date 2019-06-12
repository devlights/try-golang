package runtime_

import (
	"fmt"
	"runtime"
)

// runtime.Version() のサンプル
func Runtime01() error {
	// runtime.Version() で 現在利用している Go のバージョンが取得できる
	ver := runtime.Version()
	fmt.Println(ver)

	return nil
}
