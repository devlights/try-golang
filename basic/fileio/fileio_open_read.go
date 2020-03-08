package fileio

import (
	"fmt"
	"log"
	"os"
)

// OpenRead は、ファイルをOpenしてReadするサンプルです.
func OpenRead() error {
	// ファイルを操作する場合は os パッケージを利用する
	filename := "README.md"
	file, err := os.OpenFile(filename, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// GO言語ではIO周りのデータは文字列ではなくバイト列として扱うのみ
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// バイト列を文字列にしたい場合は string([]byte) を使う
	fmt.Printf("%s", string(buf[:n]))

	return nil
}
