package filesystem

import (
	"io/fs"
	"os"

	"github.com/devlights/gomy/output"
)

// ListDir -- os.DirFS() から fs.Glob() 経由で ディレクトリ 内のファイル一覧を出力するサンプルです.
func Listdir() error {
	cwd, err := os.Getwd()
	if err != nil {
		return nil
	}
	output.Stdoutl("[cwd ]", cwd)

	dir := os.DirFS(cwd)
	files, err := fs.Glob(dir, "go.*")
	if err != nil {
		return err
	}

	for _, f := range files {
		output.Stdoutl("[file]", f)
	}

	return nil
}
