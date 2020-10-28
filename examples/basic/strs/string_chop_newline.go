package strs

import (
	"strings"

	"github.com/devlights/gomy/output"
)

// ChopNewLine -- 文字列末尾の改行を削除するサンプルです
func ChopNewLine() error {
	var (
		withNewLine    = "helloworld\n"
		withoutNewLine = "helloworld"
	)

	output.Stdoutl("[has newline? (withNewLine)      ]", strings.HasSuffix(withNewLine, "\n"))
	output.Stdoutl("[has newline? (withoutNewLine)   ]", strings.HasSuffix(withoutNewLine, "\n"))

	// 末尾の改行を削除
	chopped := strings.TrimRight(withNewLine, "\n")

	output.Stdoutl("[has newline? (chopped)          ]", strings.HasSuffix(chopped, "\n"))
	output.Stdoutl("[equal? (chopped==withoutNewLine)]", chopped == withoutNewLine)

	output.StdoutHr()
	output.Stdoutl("[with   ]", []byte(withNewLine))
	output.Stdoutl("[chopped]", []byte(chopped))
	output.Stdoutl("[without]", []byte(withoutNewLine))

	return nil
}
