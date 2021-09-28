package debug

import (
	"errors"
	"runtime/debug"
	"strings"

	"github.com/devlights/gomy/output"
)

// BuildInfo -- debug.ReadBuildInfo を使ったサンプルです.
func BuildInfo() error {
	//
	// 実行時に依存しているライブラリの情報を調べるには runtime/debug パッケージが便利.
	// ReadBuildInfo() を使えば、ライブラリの情報が取得できる
	//
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return errors.New("debug.ReadBuildInfo() is error")
	}

	for _, dep := range info.Deps {
		if strings.Contains(dep.Path, "gomy") {
			output.Stdoutf("[dep]", "%s\t%s (%s)\n", dep.Path, dep.Version, dep.Sum)
		}
	}

	return nil
}