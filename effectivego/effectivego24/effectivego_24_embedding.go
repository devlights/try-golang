package effectivego24

import (
	"fmt"
	"log"
	"strings"

	"github.com/devlights/gomy/output"
)

type (
	helloworld struct {
	}

	embedded struct {
		name        string // 自分のフィールド
		*helloworld        // 組み込み構造体
		*log.Logger        // 組み込み構造体
	}

	embedded2 struct {
		name      string // 自分のフィールド
		*embedded        // 組み込み構造体
	}
)

func newEmbedded(h *helloworld) *embedded {
	return &embedded{
		name:       "embedded1",
		helloworld: h,
		Logger:     log.New(output.Writer(), "[embedded] ", 0),
	}
}

func newEmbedded2(e *embedded) *embedded2 {
	return &embedded2{
		name:     "embedded2",
		embedded: e,
	}
}

func (h *helloworld) pr() string {
	return "helloworld"
}

func (e *embedded) prName() string {
	h := e.helloworld.pr()
	return strings.ToUpper(h)
}

func (e *embedded2) pr() string {
	return "embedded2 pr()"
}

func (e *embedded2) prName() string {
	return fmt.Sprintf("%s -- embedded2", e.embedded.prName())
}

// Embedding -- Effective Go - Embedding の 内容についてのサンプルです。
func Embedding() error {
	/*
		https://golang.org/doc/effective_go.html#embedding

		Go には、他の言語にあるような継承の概念が無いが、型の組み込みはサポートされている。
		組み込みはインターフェースも構造体でも可能。

		構造体の方の組み込みは少しクセがあり、フィールド名を付与せずに型名のみを記載すると
		その型がまるごと組み込まれる仕組みとなっている。そのため、組み込む側の構造体で
		委譲するようなメソッド定義を行う必要がなくなる。

		当然、組み込んだ型は組み込まれた型が元々実装していたインターフェースも
		自動的に実装していることになる。

		よく利用されるシーンとして、*log.Logger を組み込みにしてしまうことで
		あたかも、その構造体がLoggerのように振る舞うことが出来る。

		組み込んでいる側のメソッドで、組み込んだ型に対してアクセスしたい場合は
		パッケージ名を除いた型名でアクセスすることが出来る。

			例： *log.Logger を組み込んでいる場合は Logger でアクセスできる

		組み込んだ型側にて予め定義されているフィールド、および、メソッドと同名の定義を
		行った場合、親側、つまり、組み込んでいる側の方が優先される。
		つまり、同名の定義を行うことで、元の定義を隠すことになる。
	*/
	e := newEmbedded(&helloworld{})

	// 組み込んだ *helloworld のメソッド
	output.Stdoutl("e.pr()", e.pr())

	// 組み込んだ *log.Logger のメソッド
	e.Println("helloworld")

	// 自身のメソッド
	e.Println(e.prName())

	// ---------------------------------------------------------------------
	// 組み込みの *helloworld が持つメソッドと同名のメソッドを定義している型で試す
	// ---------------------------------------------------------------------
	e2 := newEmbedded2(e)

	// pr() は、元々 *helloworld 側で定義されているが、親側で同名定義しているので隠される
	output.Stdoutl("e2.pr()", e2.pr())
	// prName() は、元々 *embedded 側で定義されているが、親側で同名定義しているので隠される
	output.Stdoutl("e2.prName()", e2.prName())

	// 元々のメソッドは当然存在しているので、直接指定すれば勿論呼べる
	output.Stdoutl("e2.helloworld.pr()", e2.helloworld.pr())
	output.Stdoutl("e2.embedded.prName()", e2.embedded.prName())

	return nil
}
