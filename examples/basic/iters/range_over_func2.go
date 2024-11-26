package iters

import "github.com/devlights/gomy/output"

// Go123RangeOverFunc2 は、Go 1.23 で正式導入となった Range-Over-Func のサンプルです。
//
// Range-Over-Func は、独自のイテレータを作成出来るようになる機能です。
// 以下の関数パターンがサポートされています。
//
//   - func(func() bool)     : ループ変数に値を渡さないタイプ
//   - func(func(v) bool)    : 1つのループ変数に値を渡すタイプ
//   - func(func(k, v) bool) : 2つのループ変数に値を渡すタイプ
//
// 本サンプルは、func(func(v) bool) (1つのループ変数に値を渡すタイプ) についてのサンプルです。
//
// # REFERENCES
//   - https://tip.golang.org/doc/go1.23
//   - https://tip.golang.org/blog/range-functions
//   - https://tip.golang.org/ref/spec#For_range
//   - https://pkg.go.dev/iter@go1.23.3
//   - https://zenn.dev/koya_iwamura/articles/7e7482c7222e37
//   - https://tech.every.tv/entry/2023/12/09/1
//   - https://future-architect.github.io/articles/20240129a/
func Go123RangeOverFunc2() error {
	var (
		// ２回分のイテレータ。1つのループ変数に値を渡すタイプ。
		twoTimes = func(yield func(v int) bool) {
			if !yield(100) {
				return
			}

			if !yield(99) {
				return
			}
		}
		// 指定された値からのカウントダウンを行うイテレータ。
		countdown = func(v int) func(func(int) bool) {
			return func(yield func(v int) bool) {
				for {
					if v < 0 || !yield(v) {
						return
					}

					v--
				}
			}
		}
	)

	// func(func(v) bool) のイテレータなので、ループ毎に１つのループ変数を受け取る。
	for i := range twoTimes {
		output.Stdoutf("[twoTimes    ]", "%d\n", i)
	}

	for i := range countdown(5) {
		output.Stdoutf("[countdown(5)]", "%d\n", i)
	}

	return nil
}
