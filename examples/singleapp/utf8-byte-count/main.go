package main

import (
	"flag"
	"unicode/utf8"

	"github.com/devlights/gomy/output"
)

func main() {
	var (
		u = flag.Bool("u", false, "use rune")
	)

	flag.Parse()
	if err := run(*u); err != nil {
		panic(err)
	}
}

func run(runeMode bool) error {
	var (
		strs = []string{
			// 全角かな
			"こんにちは",
			// 全角カタカナ
			"コンニチハ",
			// 半角カタカナ
			"ｺﾝﾆﾁﾜ",
			// 英数字
			"golang->60l4n6",
			// ©¼½¾
			"\U000000A9\U000000BC\U000000BD\U000000BE",
			// 🍺🍻🍷🍜
			"\U0001F37A\U0001F37B\U0001F377\U0001F35C",
		}
		fn = manual
	)

	if runeMode {
		fn = userune
		output.Stdoutl("[MODE]", "Use Rune")
	}

	for _, v := range strs {
		output.Stdoutf("", "[%s]", v)
		output.StdoutHr()
		fn(v)
	}

	return nil
}

func userune(s string) {
	//lint:ignore S1029 It's ok because this is just a example.
	//lint:ignore SA6003 It's ok because this is just a example.
	for _, r := range []rune(s) {

		if r == rune(' ') {
			output.StderrHr()
			continue
		}

		output.Stdoutl("[byte-count]", utf8.RuneLen(r))
	}
}

func manual(s string) {
	for i := 0; i < len(s); {
		//
		// UTF-8の先頭バイトを判定し、バイトサイズ算出
		//
		var (
			b = s[i]
			l = 0
		)

		switch {
		case (b & 0x80) == 0:
			l = 1
		case (b & 0xE0) == 0xC0:
			l = 2
		case (b & 0xF0) == 0xE0:
			l = 3
		case (b & 0xF8) == 0xF0:
			l = 4
		}

		i += l
		if b == ' ' {
			output.StdoutHr()
			continue
		}

		output.Stdoutl("[byte-count]", l)
	}
}
