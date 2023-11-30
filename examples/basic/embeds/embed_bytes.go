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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: embed_bytes

	   [Name] "embed_bytes"
	   00000000  68 65 6c 6c 6f 0a 77 6f  72 6c 64                 |hello.world|


	   [Elapsed] 29.61µs
	*/

}
