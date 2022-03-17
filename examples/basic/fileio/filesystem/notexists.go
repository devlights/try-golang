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
}
