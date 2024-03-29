package filepaths

import (
	"path/filepath"

	"github.com/devlights/gomy/output"
)

// FilePathGlob は、 filepath.Glob() の動作についてのサンプルです.
//
// https://golang.org/path/filepath/#Glob
func FilePathGlob() error {
	// ---------------------------------------------------------------------
	// filepath.Glob(pattern string) (matches []string, err error)
	//
	// pattern に合致するファイルをスライスにして返してくれる.
	// pattern に指定できる表記は filepath.Match() で指定できるものと同じ.
	// (https://golang.org/path/filepath/#Match)
	//
	// 注意点として、起点となるディレクトリを指定する引数はないので
	// カレント以外のディレクトリを対象としたい場合は絶対パス形式でパターンを
	// 指定する必要がある.
	//
	// 一つもファイルが見つからなかった場合は nil が返る.
	// エラーにはならないので注意.
	// ---------------------------------------------------------------------
	var (
		pattern = filepath.Join("examples", "basic", "filepaths", "*.go")
	)

	matches, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	for _, m := range matches {
		output.Stdoutl(pattern, m)
	}

	output.StdoutHr()

	// 存在しないパターンを指定
	pattern = filepath.Join("examples", "basic", "filepaths", "notexist_pattern")
	matches, err = filepath.Glob(pattern)
	if err != nil {
		return err
	}

	output.Stdoutl("matches is nil", matches == nil)
	for _, m := range matches {
		output.Stdoutl(pattern, m)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: filepath_glob

	   [Name] "filepath_glob"
	   examples/basic/filepaths/*.go examples/basic/filepaths/doc.go
	   examples/basic/filepaths/*.go examples/basic/filepaths/examples.go
	   examples/basic/filepaths/*.go examples/basic/filepaths/filepath_glob.go
	   examples/basic/filepaths/*.go examples/basic/filepaths/filepath_walk.go
	   --------------------------------------------------
	   matches is nil       true


	   [Elapsed] 152.71µs
	*/

}
