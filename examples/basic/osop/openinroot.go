package osop

import (
	"os"

	"github.com/devlights/gomy/output"
)

// OpenInRoot は、os.OpenInRoot についてのサンプルです。
//
// os.OpenInRoot(dir, name) は、以下のショートカットです。
//
//  1. os.OpenRoot(dir)
//  2. root.Open(name)
//
// # REFERENCES
//   - https://pkg.go.dev/os@go1.26.3#OpenInRoot
//   - https://pkg.go.dev/os@go1.26.3#OpenRoot
//   - https://pkg.go.dev/os@go1.26.3#Root
func OpenInRoot() error {
	var (
		cwd string
		err error
	)
	if cwd, err = os.Getwd(); err != nil {
		return err
	}

	// cwd をルートとして、ファイルを開く。
	// 配下のファイルは操作可能。
	var (
		pOk = "examples/basic/osop/openinroot.go" // このファイル
		fOk *os.File
	)
	if fOk, err = os.OpenInRoot(cwd, pOk); err != nil {
		return err
	}
	defer fOk.Close()

	// ルートの外に出ようとするとエラー
	var (
		pNg = "../" + pOk
		fNg *os.File
	)
	if fNg, err = os.OpenInRoot(cwd, pNg); err != nil {
		output.Stdoutf("[ERR]", "%v", err)
		return nil
	}
	defer fNg.Close()

	return nil
}
