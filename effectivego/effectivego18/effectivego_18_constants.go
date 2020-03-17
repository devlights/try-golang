package effectivego18

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

type ConstInt int

func (c ConstInt) String() string {
	switch {
	case c == 0:
		return fmt.Sprintf("[%d] Zero", c)
	case c == 1:
		return fmt.Sprintf("[%d] One", c)
	case c == 2:
		return fmt.Sprintf("[%d] Two", c)
	}

	return fmt.Sprintf("%d", c)
}

// 普通の int を使った定数
const (
	IntConstZero = iota
	IntConstOne
	IntConstTwo
)

// ConstInt を使った定数
const (
	ConstIntZero ConstInt = iota
	ConstIntOne
	ConstIntTwo
)

var (
	ints = []int{
		IntConstZero,
		IntConstOne,
		IntConstTwo,
	}

	constInts = []ConstInt{
		ConstIntZero,
		ConstIntOne,
		ConstIntTwo,
	}
)

func init() {
	// Go では、各ソースファイルには init() {} という初期化関数を定義することができる
	// init() {} は、複数存在しても構わない
	// init() {} は、全ての変数宣言が評価され、それらの初期化が終わり、インポートされたパッケージが
	// 初期化された後で呼び出される。

}

// Constants -- Effective Go - Constants の 内容についてのサンプルです。
func Constants() error {
	/*
		https://golang.org/doc/effective_go.html#constants

		- Goの定数は、本当にただの定数を表す
		- 値はコンパイル時に決定する必要がある
		- 可変するものは定数として利用できない
		- 連続する値を表現する場合は iota を用いて割り振るのが一般的
		- 値として利用する型に対して別名を付与して、それで定数を定義しておくのが一般的
	*/
	output.Stdoutl("(1)", IntConstOne, ConstIntOne)

	for _, c := range ints {
		output.Stdoutl("(2)", c)
	}

	for _, c := range constInts {
		output.Stdoutl("(3)", c)
	}

	return nil
}
