package strs

import (
	"strings"

	"github.com/devlights/gomy/output"
)

// StringsMap は、strings.Mapを使って全角数字を半角数字に変換するサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/strings@go1.27.3#Map
func StringsMap() error {
	//
	// strings.Mapは、変換関数を第一引数に受け取る。
	// この変換関数で、全角数字を半角数字に変換。
	//
	// 全角数字と半角数字は、それぞれ文字コードが 0〜9 の順に並んでいるため、全角の ０ からの差分を半角の 0 に加えることで変換できる。
	//
	const (
		HalfZero = '0'
		WideZero = '０'
		WideNine = '９'
	)
	var (
		s  = "12345abc１２３４５999８８８ＡＢＣ"
		fn = func(r rune) rune {
			if WideZero <= r && r <= WideNine {
				return HalfZero + (r - WideZero)
			}

			return r
		}
	)
	output.Stdoutl("[before]", s)
	output.Stdoutl("[after ]", strings.Map(fn, s))

	return nil
}
