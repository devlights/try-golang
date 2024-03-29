package stat

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Stat は、ファイル情報を取得するサンプルです.
func Stat() error {
	// ファイルの情報を取得するには、 os.Stat() を利用する
	fname := "README.md"
	fpath := filepath.Join(".", fname)

	if _, err := os.Stat(fpath); errors.Is(err, fs.ErrNotExist) {
		log.Fatal(err)
		return err
	}

	// ファイル情報取得
	fstat, err := os.Stat(fpath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// 日付のフォーマットは time.Format() を利用する
	// 予め標準的な書式は定数で定義されている
	// REF: http://bit.ly/2W4AIXh
	fmt.Printf(
		"name: %s, size: %d, date: %s, isdir: %t, mode: %s",
		fstat.Name(),
		fstat.Size(),
		fstat.ModTime().Format(time.RFC3339),
		fstat.IsDir(),
		fstat.Mode())

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: fileio_stat

	   [Name] "fileio_stat"
	   name: README.md, size: 3445, date: 2023-12-06T02:03:05Z, isdir: false, mode: -rw-r--r--

	   [Elapsed] 98.88µs
	*/

}
