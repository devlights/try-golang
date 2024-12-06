package iters

import (
	"iter"

	"github.com/devlights/gomy/output"
)

// IterPull1 は、Go 1.23 で正式導入となった iter.Pull のサンプルです。
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
// iter.Pull() または iter.Pull2() は、PULL型の挙動を行う際に利用するものです。
//
// # REFERENCES
//   - https://tip.golang.org/doc/go1.23
//   - https://tip.golang.org/blog/range-functions
//   - https://tip.golang.org/ref/spec#For_range
//   - https://pkg.go.dev/iter@go1.23.3
//   - https://zenn.dev/koya_iwamura/articles/7e7482c7222e37
//   - https://tech.every.tv/entry/2023/12/09/1
//   - https://future-architect.github.io/articles/20240129a/
func IterPull1() error {
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
		// 無理やりだけど、与えられたシーケンスを偶数or奇数のみにする処理
		pullEven = func(seq iter.Seq[int], even bool) iter.Seq[int] {
			return func(yield func(int) bool) {
				var (
					next func() (int, bool)
					stop func()
				)
				next, stop = iter.Pull(seq)
				defer stop() // 利用者側がループを完了するまえに抜けるかもしれないので必ずdeferでstop()を呼んでおく。

				var (
					v  int
					ok bool
				)
				for {
					v, ok = next()
					if !ok {
						return
					}

					ok = v%2 != 0
					if even {
						ok = v%2 == 0
					}

					if !ok {
						continue
					}

					if !yield(v) {
						return
					}
				}
			}
		}
		// 上と同じ内容をpush型で実装
		pushEven = func(seq iter.Seq[int], even bool) iter.Seq[int] {
			return func(yield func(int) bool) {
				for v := range seq {
					b := v%2 != 0
					if even {
						b = v%2 == 0
					}

					if b {
						if !yield(v) {
							return
						}
					}
				}
			}
		}
	)

	for i := range pullEven(countdown(5), true) {
		output.Stdoutf("[pullEven(5), true]", "%d\n", i)
	}

	output.StdoutHr()

	for i := range pullEven(countdown(5), false) {
		output.Stdoutf("[pullEven(5), false]", "%d\n", i)
	}

	output.StdoutHr()

	for i := range pushEven(countdown(5), true) {
		output.Stdoutf("[pushEven(5), true]", "%d\n", i)
	}

	output.StdoutHr()

	for i := range pushEven(countdown(5), false) {
		output.Stdoutf("[pushEven(5), false]", "%d\n", i)
	}

	return nil
}
