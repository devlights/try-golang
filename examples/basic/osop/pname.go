package osop

import (
	"os"
	"path/filepath"

	"github.com/devlights/gomy/output"
)

// Pname は、自身のプロセス名を取得するサンプルです。
//
// os.Executable()を利用して取得します。
//
// > Executable returns the path name for the executable that started the current process.
//
// > (Executable は、現在のプロセスを開始した実行ファイルのパス名を返します。)
//
// ただし、以下の点に注意。
//
// > There is no guarantee that the path is still pointing to the correct executable.
// If a symlink was used to start the process, depending on the operating system,
// the result might be the symlink or the path it pointed to.
//
// > (パスが正しい実行ファイルを指しているという保証はない。
// シンボリックリンクがプロセスの起動に使用された場合、オペレーティング・システムによっては、
// シンボリックリンクまたはそのシンボリックリンクが指すパスが結果として返される可能性がある。)
//
// # REFERENCES
//   - https://pkg.go.dev/os@go1.22.1#Executable
func Pname() error {
	var (
		execPath string
		err      error
	)

	execPath, err = os.Executable()
	if err != nil {
		return err
	}

	var (
		pid   = os.Getpid()
		pname = filepath.Base(execPath) // フルパスなのでファイル名のみに
	)

	output.Stdoutl("[pid  ]", pid)
	output.Stdoutl("[pname]", pname)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: osop_pname

	   [Name] "osop_pname"
	   [pid  ]              34909
	   [pname]              try-golang


	   [Elapsed] 53.71µs
	*/

}
