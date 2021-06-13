package system

import (
	"runtime"

	"github.com/devlights/gomy/output"
)

// OsName は、OSの名前を出力するサンプルです.
func OsName() error {
	output.Stdoutl("[OS]", runtime.GOOS)
	return nil
}
