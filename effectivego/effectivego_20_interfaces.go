package effectivego

import (
	"fmt"
)

type (
	sequence interface {
		copy() sequence
	}

	intSequence []int
)

func (i intSequence) copy() sequence {
	iCopy := make(intSequence, 0, len(i))
	return append(iCopy, i...)
}

func (i intSequence) String() string {
	return fmt.Sprint([]int(i))
}

// Interfaces -- Effective Go - Interfaces の 内容についてのサンプルです。
func Interfaces() error {
	/*
		https://golang.org/doc/effective_go.html#interfaces

		- Goのインターフェースは特定のオブジェクトの振る舞いを定義するためのもの
		- インターフェースを実装する際に他の言語のように明示的に「implements」などの指定が必要ない
		  - そのインターフェースが備えるメソッドを全て実装するとインターフェースを実装したという意味になる
		- Goではインターフェースはシンプルな方が良いとされている
		  - 一つのインターフェースで備えるメソッドは１つか２つくらい
		- インターフェース名は通常メソッドの名前からつけることが多い
		  - 例）Write メソッドを持つインターフェースなので Writer
		  - 例）String メソッドを持つインターフェースなので Stringer
		  - 後ろに er をつけるのが多い
		    - 他の言語だと able をつけるのが多い
	*/
	var (
		seq sequence
		fn  = func(seq sequence) {
			seqCopy := seq.copy()
			fmt.Printf("seq    : %v (%p)\nseqCopy: %v (%p)", seq, &seq, seqCopy, &seqCopy)
		}
	)

	seq = intSequence{1, 2, 3, 4, 5}
	fn(seq)

	return nil
}
