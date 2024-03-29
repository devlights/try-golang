package containers

import (
	"container/ring"

	"github.com/devlights/gomy/output"
)

// RingBuffer -- container/ring のサンプルです。（リングバッファ、循環リスト）
//
// REFERENCES
//   - https://pkg.go.dev/container/ring@go1.17.6
func RingBuffer() error {
	// container/ring は、循環リストを実装したものとなっている.
	// リング上のバッファをイメージすると分かりやすく、要素を追加していくと
	// 内容が入れ替わっていく。
	// 性質上、「直近N件分のみ保持しておきたい」場合などに便利。

	var (
		r  = ring.New(3)
		f1 = func(v interface{}) {
			output.Stdoutl("ring [next]", v)
		}
		f2 = func(v interface{}) {
			output.Stdoutl("ring [prev]", v)
		}
	)

	r.Do(f1)
	output.StdoutHr()

	for i := 0; i < 5; i++ {
		// 現在の番目に値を設定し、次の番目に進めて、それを保持しておく
		r.Value = i
		r = r.Next()

		r.Do(f1)
		output.StdoutHr()
	}

	// 循環しているのかどうかを確認
	for i := 0; i < 5; i++ {
		r = r.Prev()
		r.Do(f2)
		output.StdoutHr()
	}

	output.Stdoutl("ring-len", r.Len())

	return nil
}
