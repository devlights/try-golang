package iters

import (
	"strconv"

	"github.com/devlights/gomy/output"
)

// Go123RangeOverFunc3 は、Go 1.23 で正式導入となった Range-Over-Func のサンプルです。
//
// Range-Over-Func は、独自のイテレータを作成出来るようになる機能です。
// 以下の関数パターンがサポートされています。
//
//   - func(func() bool)     : ループ変数に値を渡さないタイプ
//   - func(func(v) bool)    : 1つのループ変数に値を渡すタイプ
//   - func(func(k, v) bool) : 2つのループ変数に値を渡すタイプ
//
// 本サンプルは、func(func(k, v) bool) (2つのループ変数に値を渡すタイプ) についてのサンプルです。
//
// # REFERENCES
//   - https://tip.golang.org/doc/go1.23
//   - https://tip.golang.org/blog/range-functions
//   - https://tip.golang.org/ref/spec#For_range
//   - https://pkg.go.dev/iter@go1.23.3
//   - https://zenn.dev/koya_iwamura/articles/7e7482c7222e37
//   - https://tech.every.tv/entry/2023/12/09/1
//   - https://future-architect.github.io/articles/20240129a/
func Go123RangeOverFunc3() error {
	var (
		// ２回分のイテレータ。2つのループ変数に値を渡すタイプ。
		twoTimes = func(yield func(i int, s string) bool) {
			if !yield(0, strconv.Itoa(100)) {
				return
			}

			if !yield(1, strconv.Itoa(99)) {
				return
			}
		}
		// 指定された文字列を逆順でループ
		reverse = func(s string) func(func(int, string) bool) {
			var (
				runes   = []rune(s)
				runeLen = len(runes)
			)

			return func(yield func(i int, s string) bool) {
				for i, j := 0, runeLen-1; i < runeLen; i, j = i+1, j-1 {
					if !yield(i, string(runes[j])) {
						return
					}
				}
			}
		}
	)

	// func(func(k, v) bool) のイテレータなので、ループ毎に2つのループ変数を受け取る。
	for i, v := range twoTimes {
		output.Stdoutf("[twoTimes           ]", "%d:%v\n", i, v)
	}

	for i, v := range reverse("helloworld") {
		output.Stdoutf("[reverse(helloworld)]", "%d:%v\n", i, v)
	}

	return nil
}
