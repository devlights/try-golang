package embeds

import (
	"embed"
	"fmt"

	"github.com/devlights/gomy/output"
)

//go:embed data
//go:embed helloworld.txt
var content embed.FS

// EmbedFsMultifiles は、embed パッケージの機能を確認するサンプルです (embed.FSとして複数ファイルを操作)
func EmbedFsMultifiles() error {
	var (
		buf []byte
		err error
	)

	//
	// 埋め込まれたファイルの操作
	//
	buf, err = content.ReadFile("helloworld.txt")
	if err != nil {
		return err
	}

	output.Stdoutl("[helloworld.txt]", string(buf))
	output.StdoutHr()

	//
	// 埋め込まれたディレクトリの操作
	//
	entries, err := content.ReadDir("data")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		var (
			name = entry.Name()
			path = fmt.Sprintf("data/%s", name)
		)

		buf, err = content.ReadFile(path)
		if err != nil {
			return err
		}

		output.Stdoutl(fmt.Sprintf("[%s]", name), string(buf))
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: embed_fs_multifiles

	   [Name] "embed_fs_multifiles"
	   [helloworld.txt]     hello
	   world
	   --------------------------------------------------
	   [dotnet.txt]         dotnet
	   [go.txt]             golang
	   [golang.txt]         golang
	   [python.txt]         python


	   [Elapsed] 64.91µs
	*/

}
