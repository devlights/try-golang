package binaries

import (
	"encoding/hex"
	"io"
	"os"
	"strings"

	"github.com/devlights/gomy/output"
)

// UsingHexDumper -- encoding/hex#Dumper のサンプルです。
func UsingHexDumper() error {
	// -----------------------------------------------------
	// encoding/hex#Dumper を利用すると hexdump コマンドを
	// 実行した場合のような16進数ダンプを出力することができる.
	//
	// hex.Dumper は、io.Writer を受け取り、io.WriteCloserを返す.
	// -----------------------------------------------------
	var (
		s      string         = "hello world"
		r      io.Reader      = strings.NewReader(s)
		w      io.Writer      = os.Stdout
		dumper io.WriteCloser = hex.Dumper(w)
	)

	output.Stdoutl("[original]", s)

	_, err := io.Copy(dumper, r)
	if err != nil {
		return err
	}

	return nil
}
