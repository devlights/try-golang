package iters

import (
	"iter"

	"github.com/devlights/gomy/output"
)

// IterSeq は、Go 1.23 で正式導入となった iter.Seq のサンプルです。
//
// Range-Over-Func は、独自のイテレータを作成出来るようになる機能です。
// 以下の関数パターンがサポートされています。
//
//   - func(func() bool)     : ループ変数に値を渡さないタイプ
//   - func(func(v) bool)    : 1つのループ変数に値を渡すタイプ
//   - func(func(k, v) bool) : 2つのループ変数に値を渡すタイプ
//
// Go 1.23 にて追加された iter パッケージには以下の２つの型が定義されています。
//
//   - iter.Seq[V any]
//   - iter.Seq2[K, V any]
//
// 上はそれぞれ func(v) bool と func(k, v) bool に対応しており、カスタムイテレータを
// 戻り値や引数として指定したりする際に、イテレータであると利用者側に伝わりやすくなります。
//
// # REFERENCES
//   - https://tip.golang.org/doc/go1.23
//   - https://tip.golang.org/blog/range-functions
//   - https://tip.golang.org/ref/spec#For_range
//   - https://pkg.go.dev/iter@go1.23.3
//   - https://zenn.dev/koya_iwamura/articles/7e7482c7222e37
//   - https://tech.every.tv/entry/2023/12/09/1
//   - https://future-architect.github.io/articles/20240129a/
func IterSeq() error {
	var (
		// 指定された値からのカウントダウンを行うイテレータ。
		countdown = func(v int) iter.Seq[int] {
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

	for i := range countdown(5) {
		output.Stdoutf("[countdown(5)]", "%d\n", i)
	}

	return nil
}
