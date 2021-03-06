package stat

import (
	"os"
	"path/filepath"

	"github.com/devlights/gomy/output"
)

// StatPermission は、ファイルのパーミッションに関するサンプルです。
func StatPermission() error {
	const fpath = "examples/basic/fileio/fileio_permission.go"

	absPath, err := filepath.Abs(fpath)
	if err != nil {
		return err
	}
	output.Stdoutl("[abspath]", absPath)
	output.StdoutHr()

	fstat, err := os.Stat(absPath)
	if err != nil {
		return err
	}

	var fmode os.FileMode = fstat.Mode()
	var fperm os.FileMode = fmode.Perm()

	if fperm&0400 == 0400 {
		output.Stdoutl("[readable]", "yes")
	}

	if fperm&0200 == 0200 {
		output.Stdoutl("[writable]", "yes")
	}

	if fperm&0100 == 0100 {
		output.Stdoutl("[executable]", "yes")
	}

	output.StdoutHr()
	output.Stdoutf("[Permission]", "%[1]v\t%#[1]o\n", fperm)

	return nil
}
