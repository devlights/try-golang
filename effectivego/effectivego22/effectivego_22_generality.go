package effectivego22

import (
	"github.com/devlights/gomy/output"
)

type (
	runner interface {
		run()
	}

	setter interface {
		set(val string)
	}

	helloworldRunner struct {
		value string
	}
	hogeRunner struct {
		value string
	}
)

func newHelloworldRunner(val string) runner {
	return &helloworldRunner{
		value: val,
	}
}

func newHogeRunner(val string) runner {
	return &hogeRunner{
		value: val,
	}
}

func (h *helloworldRunner) run() {
	output.Stdoutl("helloworldRunner", h.value)
}

func (h *hogeRunner) run() {
	output.Stdoutl("hogeRunner", h.value)
}

func (h *helloworldRunner) set(val string) {
	h.value = val
}

func (h *hogeRunner) set(val string) {
	h.value = val
}

// Generality -- Effective Go - Generality の 内容についてのサンプルです。
func Generality() error {
	/*
		https://golang.org/doc/effective_go.html#generality

		特定のインターフェースを実装しているのみの型が存在する場合
		その型を外部に公開する必要は特に無い。

		このような場合は、型自体を非公開、つまり「先頭小文字で始まる名前」で定義し
		コンストラクタ関数にてインターフェースを返すようにしておくと、インターフェースが
		持つ振る舞い以上の動作を持っていないことが明確となる。
	*/
	var (
		r1 = newHelloworldRunner("hello")
		r2 = newHogeRunner("ho")
	)

	r1.run()
	r2.run()

	// 別のインターフェースとしても利用したい場合は　val.(typeName) の 形式で変換する
	if v, ok := r1.(setter); ok {
		v.set("world")
	}

	if v, ok := r2.(setter); ok {
		v.set("hoge")
	}

	r1.run()
	r2.run()

	return nil
}
