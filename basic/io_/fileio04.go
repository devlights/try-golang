package io_

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func FileIo04() error {
	// ファイルの情報を取得するには、 os.Stat() を利用する
	fname := "README.md"
	fpath := filepath.Join(".", fname)

	if _, err := os.Stat(fpath); os.IsNotExist(err) {
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
}
