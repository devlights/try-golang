package filesystem

import (
	"os"
	"strings"

	"github.com/devlights/gomy/output"
)

// ReadDir は、 os.ReadDir() を使ったサンプルです.
func ReadDir() error {
	var (
		pwd string
		err error
	)

	if pwd, err = os.Getwd(); err != nil {
		return err
	}

	//
	// os.ReadDir() では、fs.FS ではなく、os.DirEntry が
	// 返却されることに注意.
	//
	var (
		entries []os.DirEntry
	)

	if entries, err = os.ReadDir(pwd); err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		dirname := entry.Name()
		if strings.HasPrefix(dirname, ".") {
			continue
		}

		output.Stdoutl("[directory]", dirname)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: fileio_filesystem_readdir

	   [Name] "fileio_filesystem_readdir"
	   [directory]          builder
	   [directory]          cmd
	   [directory]          examples
	   [directory]          mapping
	   [directory]          runner


	   [Elapsed] 179.57µs
	*/

}
