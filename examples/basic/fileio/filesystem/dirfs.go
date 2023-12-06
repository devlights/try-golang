package filesystem

import (
	"io"
	"io/fs"
	"os"

	"github.com/devlights/gomy/output"
)

// DirFS は、 os.DirFS() を利用してファイルの読み込みを行うサンプルです.
//
// REFERENCES:
//   - https://golang.org/io/fs/
//   - https://golang.org/os/#DirFS
func DirFS() error {
	//
	// io/fs パッケージは Go1.16で追加されたパッケージ.
	// このパッケージの目的はファイルシステムを抽象化するためにある.
	//
	// 今までの io.Reader や io.Writer が読み書きを抽象化していたのと
	// 同様にファイルシステムという概念も抽象化してしまおうという事になる.
	//
	// osパッケージ側には DirFS という関数が追加され、ここから特定のディレクトリの
	// fs.FS を取得することが出来るようになった。
	//

	//
	// 以下は os.Open("./.gitpod.yml") しているのと同じことになる
	//
	var (
		dir  fs.FS
		file fs.File
		buf  []byte
		err  error
	)

	// DirFSを利用して、特定のディレクトリを指す fs.FS を取得
	dir = os.DirFS(".")
	output.Stdoutf("[DirFS]", "%[1]v(%[1]T)\n", dir)

	// 対象の fs.FS から、そのディレクトリ内のファイルを開く
	file, err = dir.Open(".gitpod.yml")
	if err != nil {
		return err
	}

	buf, err = io.ReadAll(file)
	if err != nil {
		return nil
	}

	output.Stdoutl("[.gitpod.yml]", string(buf))

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: fileio_filesystem_dirfs

	   [Name] "fileio_filesystem_dirfs"
	   [DirFS]              .(os.dirFS)
	   [.gitpod.yml]        image:
	     file: .gitpod.Dockerfile

	   tasks:
	     - name: initial
	       init:
	         go install github.com/go-task/task/v3/cmd/task@latest &&
	         go install honnef.co/go/tools/cmd/staticcheck@latest &&
	         go install golang.org/x/tools/cmd/goimports@latest &&
	         go install github.com/mgechev/revive@latest &&
	         go install github.com/go-delve/delve/cmd/dlv@latest &&
	         go install go.uber.org/nilaway/cmd/nilaway@latest &&
	         task build
	       command:
	         go version


	   vscode:
	     extensions:
	       - golang.go



	   [Elapsed] 97.109µs
	*/

}
