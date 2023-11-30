package embeds

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/devlights/gomy/output"
)

//go:embed data/go*.txt
var filteredContent embed.FS // 埋め込むファイルをフィルタリングしている

// EmbedFsFilter は、embed パッケージの機能を確認するサンプルです (埋め込むファイルをフィルタリング)
func EmbedFsFilter() error {
	var (
		entries []fs.DirEntry
		buf     []byte
		err     error
	)

	entries, err = filteredContent.ReadDir("data")
	if err != nil {
		return err
	}

	for _, entry := range entries {
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

	   ENTER EXAMPLE NAME: embed_fs_filter

	   [Name] "embed_fs_filter"
	   [go.txt]             golang
	   [golang.txt]         golang


	   [Elapsed] 35.84µs
	*/

}
