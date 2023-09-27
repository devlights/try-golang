package main

import (
	"flag"
	"fmt"
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
			// 全角かな (3bytes)
			"こんにちは",
			// 全角カタカナ (3bytes)
			"コンニチハ",
			// 半角カタカナ (3bytes)
			"ｺﾝﾆﾁﾊ",
			// 英数字記号 (1byte)
			"golang->60l4n6",
			// ©¼½¾ (2bytes)
			"\U000000A9\U000000BC\U000000BD\U000000BE",
			// 🍺🍻🍷🍜 (4bytes)
			"\U0001F37A\U0001F37B\U0001F377\U0001F35C",
		}
		fn = manual
	)

	if runeMode {
		fn = useRune
	}

	for _, v := range strs {
		output.Stdoutf("", "[%s]", v)
		output.StdoutHr()

		if err := fn(v); err != nil {
			return err
		}
	}

	return nil
}

func manual(s string) error {

	for i := 0; i < len(s); {
		var (
			b = s[i]
			l = 0
		)

		//
		// UTF-8の先頭バイトを判定し、バイトサイズ算出
		//
		// UTF-8エンコーディングでは、各文字は1バイトから4バイトまでの可変長でエンコードされる。
		// 先頭バイト（最初のバイト）を見ることで、その文字が何バイトでエンコードされているかを判定できる。
		//
		// - 0xxxxxxx: 1バイト（ASCIIと互換性あり）
		// - 110xxxxx: 続く1バイトと合わせて2バイト
		// - 1110xxxx: 続く2バイトと合わせて3バイト
		// - 11110xxx: 続く3バイトと合わせて4バイト
		//
		// 以下の case は上記を判定している.
		//
		// - (b & 0x80) == 0   : 最上位1ビットが0    であるなら、この文字は1バイト
		// - (b & 0xE0) == 0xC0: 最上位2ビットが110  であるなら、この文字は2バイト
		// - (b & 0xF0) == 0xE0: 最上位3ビットが1110 であるなら、この文字は3バイト
		// - (b & 0xF8) == 0xF0: 最上位4ビットが11110であるなら、この文字は4バイト
		//
		// REFERENCES:
		//   - https://ja.wikipedia.org/wiki/UTF-8
		//
		switch {
		case (b & 0x80) == 0:
			l = 1
		case (b & 0xE0) == 0xC0:
			l = 2
		case (b & 0xF0) == 0xE0:
			l = 3
		case (b & 0xF8) == 0xF0:
			l = 4
		default:
			return fmt.Errorf("invalid utf-8 char (%b)", b)
		}

		output.Stdoutf("[byte-count]", "%s (%d)\n", s[i:i+l], l)

		i += l
	}

	return nil
}

func useRune(s string) error {

	for _, r := range s {
		l := utf8.RuneLen(r)
		if l == -1 {
			return fmt.Errorf("invalid utf-8 char (%c)", r)
		}

		output.Stdoutf("[byte-count]", "%c (%d)\n", r, l)
	}

	return nil
}
