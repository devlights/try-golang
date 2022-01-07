package hexop

import (
	"encoding/hex"
	"io"
	"os"
	"strings"
)

// Dumper -- encoding/hex#Dumper のサンプルです。
func Dumper() error {
	// -----------------------------------------------------
	// encoding/hex#Dumper を利用すると hexdump コマンドを
	// 実行した場合のような16進数ダンプを出力することができる.
	//
	// hex.Dumper は、io.Writer を受け取り、io.WriteCloserを返す.
	//
	// 最後に Close を呼ばないと、ダンプ出力の右側に元の値が表示されないので注意
	// (Close を呼ばないままだと、16進部分のみが出力される)
	// -----------------------------------------------------
	var (
		reader = strings.NewReader("hello world")
		writer = os.Stdout
		dumper = hex.Dumper(writer)
	)
	defer dumper.Close() // Close を呼ぶことにより、出力の右側に値が出力される

	if _, err := io.Copy(dumper, reader); err != nil {
		return err
	}

	return nil
}
