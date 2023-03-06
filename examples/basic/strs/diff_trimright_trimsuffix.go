package strs

import (
	"strings"

	"github.com/devlights/gomy/output"
)

// DiffTrimRightAndTrimSuffix は、strings.TrimRight と strings.TrimSuffix のちょっとした違いについてのサンプルです.
//
// TrimRight は、第二引数が cutset となっている通り、切り取る対象の「セット」を指定している。
// なので、セット内に存在する文字であるかどうかで判定される。
//
// TrimSuffix は、第二引数が suffix となっている通り、指定した suffix に一致するかどうかが判定される。
// なので、指定した suffix に完全一致しないと除去されない。
//
// # REFERENCES
//   - https://pkg.go.dev/strings@go1.20.1#TrimRight
//   - https://pkg.go.dev/strings@go1.20.1#TrimSuffix
func DiffTrimRightAndTrimSuffix() error {
	var (
		str = "こんにちわ世界xox"
	)

	output.Stdoutl("[TrimRight ]", strings.TrimRight(str, "xo"))
	output.Stdoutl("[TrimSuffix]", strings.TrimSuffix(str, "xo"))

	return nil
}
