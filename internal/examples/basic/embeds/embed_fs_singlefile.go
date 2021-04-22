package embeds

import (
	"embed"
	"encoding/hex"
	"os"
)

//go:embed helloworld.txt
var f embed.FS

// EmbedFsSingleFile は、embed パッケージの機能を確認するサンプルです (embed.FSとして単独ファイルを操作)
func EmbedFsSingleFile() error {
	b, err := f.ReadFile("helloworld.txt")
	if err != nil {
		return err
	}

	dumper := hex.Dumper(os.Stdout)
	defer dumper.Close()

	if _, err := dumper.Write(b); err != nil {
		return err
	}

	return nil
}
