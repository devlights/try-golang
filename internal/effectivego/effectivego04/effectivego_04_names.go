package effectivego04

import "fmt"

type (
	point struct {
		x, y int
	}
)

// X is a Getter of type point.x.
func (p *point) X() int {
	return p.x
}

// SetX is a Setter of type point.x.
func (p *point) SetX(x int) {
	p.x = x
}

// Names -- Effective Go - Names の 内容についてのサンプルです。
func Names() error {
	/*
		https://golang.org/doc/effective_go.html#names

		- Go では、先頭を大文字にすることで外部公開のスコープになる
		- パッケージの名前は 小文字で単語 の形が望ましい。アンダースコアなどは出来れば付与しない。
		- Go では、言語として自動的なGetterとSetterの仕組みはサポートしていない
		  - Getter の名前は、 フィールド名の先頭を大文字にしたものが望ましい。 (e.g. owner -> Owner)
		  - Setter の名前は、 Getter名の前にSetを付与したものが望ましい。 (e.g. owner -> SetOwner)
		- インターフェースの名前として、ただ一つのメソッドを持つインターフェースの場合は 名前er とした形で定義する。
		  - e.g. Reader, Writer, Formatter, Stringer etc.
		- Go では、名前付けに MixedCaps or mixedCaps 形式を使う。
	*/

	p := &point{1, 2}
	fmt.Printf("p.x [%d]\n", p.X())

	p.SetX(99)
	fmt.Printf("p.x [%d]\n", p.X())

	return nil
}
