package filepaths

import (
	"path/filepath"
	"strings"

	"github.com/devlights/gomy/output"
)

// ExcludeSuffix は、ファイル名から拡張子を除いた値を取得するサンプルです.
//
// filepath.Ext()とstrings.TrimSuffix()を組み合わせて取得出来ます。
//
// > strings.TrimSuffix(file, filepath.Ext(file))
//
// または、スライスを拡張子分だけカットすることでも取得出来ます。
//
// > file[:len(file)-len(filepath.Ext(file))]
//
// # REFERENCES
//   - https://pkg.go.dev/path/filepath@go1.23.4#Ext
//   - https://pkg.go.dev/strings@go1.23.4#TrimSuffix
func ExcludeSuffix() error {
	var (
		fpath = "/path/to/src/something.go"
		dir   = filepath.Dir(fpath)
		fname = filepath.Base(fpath)
		ext   = filepath.Ext(fname)

		base1 = strings.TrimSuffix(fname, ext)
		base2 = fname[:len(fname)-len(ext)]
	)
	output.Stdoutl("[fpath]", fpath)
	output.Stdoutl("[dir  ]", dir)
	output.Stdoutl("[ext  ]", ext)
	output.Stdoutl("[base1]", base1)
	output.Stdoutl("[base2]", base2)

	return nil
}
