package formatting

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

// Append -- Go 1.19 から追加された fmt.Append() のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/fmt@go1.19#Append
//   - https://dev.to/emreodabas/quick-guide-go-119-features-1j40
func Append() error {
	var (
		buf = []byte("hello")
		sli = []any{"world", "hello"}
	)

	buf = fmt.Append(buf, " ")
	buf = fmt.Append(buf, "world")
	buf = fmt.Append(buf, 12345)
	buf = fmt.Append(buf, sli) // スライスをそのまま渡すと %s した状態になるので注意
	buf = fmt.Append(buf, sli...)

	output.Stdoutl("[fmt.Append]", string(buf))

	return nil
}
