package ifs

import "fmt"

type (
	// 普通のインターフェース１
	if01 interface {
		f1() string
	}

	//　普通のインターフェース２
	if02 interface {
		f2() string
	}

	// 合成インターフェース. if01とif02というインターフェースを合成している.
	// このインターフェースをするには if01とif02のインターフェース定義を満たす必要がある
	ifComposition interface {
		if01
		if02
	}

	// 実装
	compositionImpl struct{}
)

// impl: if01.f1
func (c *compositionImpl) f1() string {
	return "hello"
}

// impl: if02.f2
func (c *compositionImpl) f2() string {
	return "world"
}

// Composition は、 Goのインターフェースのコンポジション (合成)　についてのサンプルです.
func Composition() error {
	// ----------------------------------------------------------------
	// インターフェースのコンポジションについて
	//
	// Goでは、インターフェースを合成する場合、合成インターフェースを定義して
	// 属性に各インターフェースを列挙していく.
	//
	// io.ReadCloser, io.ReadWriter などが合成インターフェースの代表例
	// ----------------------------------------------------------------
	var (
		c  = &compositionImpl{}
		f1 = func(i if01) {
			fmt.Println(i.f1())
		}
		f2 = func(i if02) {
			fmt.Println(i.f2())
		}
		f3 = func(i ifComposition) {
			fmt.Println(i.f1(), i.f2())
		}
	)

	// 具象型からインターフェースへ
	var v ifComposition = c

	// 合成インターフェースは、合成元となっている各インターフェースを名乗ることが出来る
	f1(v)
	f2(v)
	f3(v)

	return nil
}
