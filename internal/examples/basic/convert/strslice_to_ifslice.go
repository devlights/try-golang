package convert

import (
	"github.com/devlights/gomy/output"
)

// StringSliceToInterfaceSlice -- []string から []interface{} への変換についてのサンプルです.
func StringSliceToInterfaceSlice() error {
	// -------------------------------------------------------------------------------
	// []string から []interface{} への変換
	//
	// Go は、型についてとても厳しい言語.
	// なので、[]string から []interface{} への変換は、直接出来ない.
	// というか、上記の２つの型はメモリ上で同じ表現とならない。
	//
	// 変換したい場合は、ユーザが明示的に変換元のスライスをループさせて
	// 要素を一つずつ コピー していく
	//
	// REFERENCES:
	//   - https://golang.org/doc/faq#convert_slice_of_interface
	//   - https://golang.org/doc/faq#convert_slice_with_same_underlying_type
	//   - https://stackoverflow.com/questions/27689058/convert-string-to-interface
	//   - https://stackoverflow.com/questions/23148812/whats-the-meaning-of-interface/23148998#23148998
	//   - https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
	//   - https://research.swtch.com/interfaces
	// -------------------------------------------------------------------------------
	var (
		// string の スライス
		strSlice = []string{
			"hello",
			"world",
		}

		// interface{} の スライス
		ifSlice []interface{}
	)

	// 以下はコンパイルできない
	//   cannot use strSlice (type []string) as type []interface {} in assignment
	// ifSlice = strSlice

	// Go では以下のように明示的にコピーしていく方法をとる
	for _, s := range strSlice {
		ifSlice = append(ifSlice, s)
	}

	output.Stdoutl("[ifSlice]", ifSlice)

	// なので同様に 他のスライス を変換する場合でも同じ理由でコンパイルエラーとなる
	var (
		byteSlice = []byte{1, 2}
		intSlice  []int
	)

	// 以下はコンパイルできない
	//   cannot use byteSlice (type []byte) as type []int in assignment
	// intSlice = byteSlice

	for _, b := range byteSlice {
		intSlice = append(intSlice, int(b))
	}

	output.Stdoutl("[intSlice]", intSlice)

	// -------------------------------------------------------------------------------
	// 以下、Go FAQ の
	//   Can I convert []T1 to []T2 if T1 and T2 have the same underlying type?
	//   https://golang.org/doc/faq#convert_slice_with_same_underlying_type
	// の内容。
	//
	// Goでは、型はメソッドと密接に結びついており、名前のついた型には（空の可能性がある）メソッドセットがあります。
	// 一般的なルールとしては、変換する型の名前は変更できますが、
	// 複合型の要素の名前（およびメソッドセット）を変更することはできません。
	// Goでは、型の変換を明示的に行う必要があります。
	//
	// 複合型というのは 構造体, 配列, スライス, マップ のこと
	// 言語仕様により変換を **明示的に**　行う必要があるため、そのまま変換できない

	// 「変換する型の名前は変更できる」というのは以下の意味
	var (
		s = "helloworld" // string 型
		i interface{}    // interface{} 型
	)

	// これはOK -- s は
	i = s
	output.Stdoutl("[s --> i]", i)

	type (
		myInt1 int
		myInt2 int
	)

	var (
		i1 myInt1
		i2 myInt2
	)

	// これはコンパイルエラー
	// 同じ型からのエイリアスであるが、ランタイムから見ると明確に異なる型として認識している
	// i2 = i1

	// これはOK -- 型の名前を変更している. つまり型変換している.
	i2 = myInt2(i1)
	output.Stdoutl("[i2]", i2)

	var is1 = []myInt1{1, 2, 3, 4}
	var is2 []myInt2

	// これはNG -- 複合型の要素の名前は変更できない. 変換を明示的に行う必要がある.
	// is2 = ([]myInt2)(is1)

	// 変換を明示的に行えばOK
	for _, v := range is1 {
		is2 = append(is2, myInt2(v))
	}

	output.Stdoutl("[is2]", is2)

	return nil
}
