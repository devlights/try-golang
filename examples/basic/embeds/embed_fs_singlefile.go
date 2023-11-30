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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: embed_fs_singlefile

	   [Name] "embed_fs_singlefile"
	   00000000  68 65 6c 6c 6f 0a 77 6f  72 6c 64                 |hello.world|


	   [Elapsed] 45.71µs
	*/

}
