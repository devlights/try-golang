package filesystem

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/devlights/gomy/output"
)

// NotExists -- ファイルの存在チェックを行うサンプルです.
//
// REFERENCES
//   - https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
func NotExists() error {
	var (
		cwd string
		err error
	)

	cwd, err = os.Getwd()
	if err != nil {
		return err
	}

	var (
		path = filepath.Join(cwd, "not_exists_filename")
	)

	if _, err = os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		output.Stdoutf("[not found]", "%s\t(%v)", path, err)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: fileio_filesystem_notexists

	   [Name] "fileio_filesystem_notexists"
	   [not found]          /workspace/try-golang/not_exists_filename  (stat /workspace/try-golang/not_exists_filename: no such file or directory)

	   [Elapsed] 68.78µs
	*/

}
