package embeds

import (
	_ "embed"
	"encoding/hex"
	"os"
)

//go:embed helloworld.txt
var b []byte

// EmbedBytes は、embed パッケージの機能を確認するサンプルです (バイト列として値を取得)
func EmbedBytes() error {
	dumper := hex.Dumper(os.Stdout)
	defer dumper.Close()

	if _, err := dumper.Write(b); err != nil {
		return err
	}

	return nil
}
