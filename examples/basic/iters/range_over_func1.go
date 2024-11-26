package iters

import "github.com/devlights/gomy/output"

// Go123RangeOverFunc1 は、Go 1.23 で正式導入となった Range-Over-Func のサンプルです。
//
// Range-Over-Func は、独自のイテレータを作成出来るようになる機能です。
// 以下の関数パターンがサポートされています。
//
//   - func(func() bool)     : ループ変数に値を渡さないタイプ
//   - func(func(v) bool)    : 1つのループ変数に値を渡すタイプ
//   - func(func(k, v) bool) : 2つのループ変数に値を渡すタイプ
//
// 本サンプルは、func(func() bool) (ループ変数に値を渡さないタイプ) についてのサンプルです。
//
// # REFERENCES
//   - https://tip.golang.org/doc/go1.23
//   - https://tip.golang.org/blog/range-functions
//   - https://tip.golang.org/ref/spec#For_range
//   - https://pkg.go.dev/iter@go1.23.3
//   - https://zenn.dev/koya_iwamura/articles/7e7482c7222e37
//   - https://tech.every.tv/entry/2023/12/09/1
//   - https://future-architect.github.io/articles/20240129a/
func Go123RangeOverFunc1() error {
	var (
		// ２回分のイテレータ。ループ変数に値を渡さないタイプ。
		twoTimes = func(yield func() bool) {
			if !yield() {
				return
			}

			if !yield() {
				return
			}
		}
		// クロージャを使って、指定された回数ループするイテレータを返す。
		nTimes = func(times int) func(func() bool) {
			return func(yield func() bool) {
				for range times {
					if !yield() {
						return
					}
				}
			}
		}
	)

	// func() bool のイテレータなので、ループ毎に受け取るものが無い。
	// (range over twoTimes (variable of type func(yield func() bool)) permits no iteration variables)
	for range twoTimes {
		output.Stdoutl("[twoTimes ]", "call")
	}

	for range nTimes(5) {
		output.Stdoutl("[nTimes(5)]", "call")
	}

	return nil
}
