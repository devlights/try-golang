package strs

import (
	"strings"

	"github.com/devlights/gomy/output"
)

// ChopNewLine -- 文字列末尾の改行を削除するサンプルです
func ChopNewLine() error {
	var (
		withLF         = "helloworld\n"
		withCRLF       = "helloworld\r\n"
		withoutNewLine = "helloworld"
	)

	output.Stdoutl("[has newline? (withLF)           ]", strings.HasSuffix(withLF, "\n"))
	output.Stdoutl("[has newline? (withCRLF)         ]", strings.HasSuffix(withCRLF, "\n"))
	output.Stdoutl("[has newline? (withoutNewLine)   ]", strings.HasSuffix(withoutNewLine, "\n"))

	// 末尾の改行を削除
	chopLF := chop(withLF)
	output.Stdoutl("[has newline? (chopped)(LF)      ]", strings.HasSuffix(chopLF, "\n"))
	output.Stdoutl("[equal? (chopped==withoutNewLine)]", chopLF == withoutNewLine)

	chopCRLF := chop(withCRLF)
	output.Stdoutl("[has newline? (chopped)(CRLF)    ]", strings.HasSuffix(chopCRLF, "\n"))
	output.Stdoutl("[equal? (chopped==withoutNewLine)]", chopCRLF == withoutNewLine)

	output.StdoutHr()
	output.Stdoutl("[LF      ]", []byte(withLF))
	output.Stdoutl("[CRLF    ]", []byte(withCRLF))
	output.Stdoutl("[chopLF  ]", []byte(chopLF))
	output.Stdoutl("[chopCRLF]", []byte(chopCRLF))
	output.Stdoutl("[without ]", []byte(withoutNewLine))

	return nil
}

func chop(s string) string {
	s = strings.TrimRight(s, "\n")
	if strings.HasSuffix(s, "\r") {
		s = strings.TrimRight(s, "\r")
	}

	return s
}
