// プラグインとして利用する処理は main パッケージに属している必要がある

package main

import (
	"fmt"

	"github.com/devlights/try-golang/examples/singleapp/073.stdlib_plugin_pkg/lib/pkg/strs"
)

func Fn(message string) {
	fmt.Println(strs.Upper(message))
}
