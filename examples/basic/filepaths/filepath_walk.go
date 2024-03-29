package filepaths

import (
	"os"
	"path/filepath"

	"github.com/devlights/gomy/output"
)

// FilePathWalk は、filepaths.Walk() のサンプルです.
func FilePathWalk() error {
	// -----------------------------------------------------
	// filepath.Walk() は、指定された起点ディレクトリから再帰的に
	// ファイルを処理していってくれる関数。
	//
	// Python の os.walk() と同じようなイメージ
	//
	// 第一引数に起点、第２引数にWalkFunc型の関数を指定する.
	// ファイルツリーを下る処理は filepath.Walk() 内で実施され
	// ファイル毎に引数に指定した WalkFunc が呼ばれる。
	// -----------------------------------------------------
	err := filepath.Walk("examples/basic/filepaths", walkFn)
	if err != nil {
		return err
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: filepath_walk

	   [Name] "filepath_walk"
	   ---------------------------------
	   path                 examples/basic/filepaths
	   isdir                true
	   info                 filepaths
	   ---------------------------------
	   ---------------------------------
	   path                 examples/basic/filepaths/README.md
	   isdir                false
	   info                 README.md
	   ---------------------------------
	   ---------------------------------
	   path                 examples/basic/filepaths/doc.go
	   isdir                false
	   info                 doc.go
	   ---------------------------------
	   ---------------------------------
	   path                 examples/basic/filepaths/examples.go
	   isdir                false
	   info                 examples.go
	   ---------------------------------
	   ---------------------------------
	   path                 examples/basic/filepaths/filepath_glob.go
	   isdir                false
	   info                 filepath_glob.go
	   ---------------------------------
	   ---------------------------------
	   path                 examples/basic/filepaths/filepath_walk.go
	   isdir                false
	   info                 filepath_walk.go
	   ---------------------------------


	   [Elapsed] 450.67µs
	*/

}

func walkFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	output.Stdoutl("---------------------------------")
	defer output.Stdoutl("---------------------------------")

	// path は、起点からの相対パスが入っている
	// info.Name() は、現在処理対象のファイルの名前が取れる
	//
	// なので、絶対パスが欲しい場合は filepath.Abs(path) とすれば良い
	output.Stdoutl("path", path)
	output.Stdoutl("isdir", info.IsDir())
	output.Stdoutl("info", info.Name())

	// absPath, _ := filepath.Abs(path)
	// output.Stdoutl("abspath", absPath)

	return nil
}
