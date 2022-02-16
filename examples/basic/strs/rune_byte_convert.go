package strs

import (
	"fmt"
	"unicode/utf8"
)

// RuneByteConvert は、文字列とルーンとバイト列の変換のサンプルです.
func RuneByteConvert() error {
	// -----------------------------------------------------
	// Goでの文字列は string 型で表現される。
	// string は、文字列をバイト列で表現しているもの。
	// なので、string に対して len() とするとバイト数が取得できる
	// (Goの文字列表現はUTF-8)
	//
	// Goでの文字は rune 型で表現される。
	// rune は、文字をUnicodeコードポイントで表現しているもの。
	// なので、 []rune に対して len() とすると文字数が取得できる。
	//
	// string/[]rune/[]byte は相互に変換可能。
	//
	// []byte に string を与える事によって、その文字列のバイト列になる
	// string に []byte を与える事によって、バイト列の文字列となる
	// []rune に string を与える事によって、その文字列のrune列となる
	// string に []rune を与える事によって、そのrune列の文字列となる
	//
	// 上を見る通り、stirng が中継地点となっていて
	// rune と byte を変換する場合は、一旦 string にしてから
	// 望む方向へ変換する。
	//
	// utf8.EncodeRune() を利用することで中間生成のstringを省く事もできる.
	//
	// REFERENCES::
	//   - https://stackoverflow.com/questions/29255746/how-encode-rune-into-byte-using-utf8-in-golang
	//   - https://qiita.com/masakielastic/items/01a4fb691c572dd71a19
	//   - https://golang.org/ref/spec#Conversions_to_and_from_a_string_type
	// -----------------------------------------------------
	s := "こんにちわworld"

	// そのまま len(s) とすると、バイト数となる
	fmt.Printf("len(s) == %d\n", len(s))

	// runeに変換
	r := []rune(s)
	fmt.Printf("len(r) == %d\n", len(r))

	// runeはUnicodeコードポイントを示しているので日本語などでは
	// 一つのruneがNバイトとなる.
	for i, v := range r {
		// runeをバイト列に変換
		// 一旦、runeを文字列にして、そこからバイト列に変換する
		b := []byte(string(v))

		fmt.Printf("rune[%d] %d byte(s)\n", i, len(b))
	}

	// ----------------------------------------------------------
	// runeのスライスを string を経由せずに、直接 byte スライスに変換
	// (https://stackoverflow.com/questions/29255746/how-encode-rune-into-byte-using-utf8-in-golang)
	// ----------------------------------------------------------
	// 予めUtf-8の１文字での最大バイト数分で大きめにバッファを用意しておく
	buf := make([]byte, len(r)*utf8.UTFMax)

	// rune を一つずつエンコード
	count := 0
	for _, v := range r {
		count += utf8.EncodeRune(buf[count:], v)
	}

	// スライスのサイズを調整
	buf = buf[:count]

	fmt.Printf("%s(%v)\n", buf, buf)

	return nil
}
