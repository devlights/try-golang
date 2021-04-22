package embeds

import (
	"embed"
	"fmt"

	"github.com/devlights/gomy/output"
)

//go:embed data
//go:embed helloworld.txt
var content embed.FS

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
	dirs, err := content.ReadDir("data")
	if err != nil {
		return err
	}

	for _, entry := range dirs {
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
}
