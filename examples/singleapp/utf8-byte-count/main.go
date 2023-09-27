package main

import (
	"unicode/utf8"

	"github.com/devlights/gomy/output"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var (
		s = "こんにちは コンニチハ ｺﾝﾆﾁﾜ hello"
	)

	manual(s)
	output.StdoutHr()
	userune(s)

	return nil
}

func userune(s string) {
	//lint:ignore S1029 It's ok because this is just a example.
	//lint:ignore SA6003 It's ok because this is just a example.
	for _, r := range []rune(s) {
		l := utf8.RuneLen(r)

		if r == rune(' ') {
			output.StderrHr()
		} else {
			output.Stdoutl("[userune][byte-count]", l)
		}
	}
}

func manual(s string) {
	for i := 0; i < len(s); {
		c := s[i]

		//
		// UTF-8の先頭バイトを判定し、バイトサイズ算出
		//
		l := 0
		switch {
		case (c & 0x80) == 0:
			l = 1
		case (c & 0xE0) == 0xC0:
			l = 2
		case (c & 0xF0) == 0xE0:
			l = 3
		case (c & 0xF8) == 0xF0:
			l = 4
		}

		if c == ' ' {
			output.StdoutHr()
		} else {
			output.Stdoutl("[manual][byte-count]", l)
		}

		i += l
	}
}
