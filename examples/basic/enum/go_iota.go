package enum

import (
	"github.com/devlights/gomy/output"
)

// GoIota -- Go における iota の扱い方についてのサンプルです
//
// REFERENCES:
//   -https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3
//
//noinspection GoBoolExpressions
func GoIota() error {
	// ---------------------------------------------------------------------------
	// iota は Goで利用できる数値のカウンタ.
	// 以下の基本的な特徴を持つ.
	//
	// - 0から始まる
	// - 呼び出すたびに1ずつ加算していく
	// - const定義の中でのみ利用できる
	// - リセットできる
	//
	// ---------------------------------------------------------------------------
	type (
		Weekday int
	)

	// 基本的な使い方
	const (
		_      Weekday = iota // 初期値は 0. 曜日は 1 からスタートしたいので捨てる
		Sunday                // = iota を省略可能。値は1となる.
		_                     // iota の値は加算されるが捨てる。つまり2を捨てる
		// コメントの場合 iota は加算をスキップ。なので増えない。空行の場合も増えない

		Monday  // 3
		Tuesday // 4
	)

	output.Stdoutl("Weekday", Sunday, Monday, Tuesday)

	// iota は const ブロック 毎にリセットされる
	const (
		Sunday2 Weekday = iota // 0
		Monday2         = iota // 1
	)

	output.Stdoutl("Weekday2", Sunday2, Monday2)

	// iota を 途中で使うと、そこまで毎行 iota を使っていた状態の値を返す
	// つまり、以下の場合だと Three の部分で iota を書くと 2 となる
	// (One の部分で 0 -> 1, Two の部分で 1 -> 2 だから)
	const (
		One   = 1        // iota 0 -> 1
		Two   = 2        // iota 1 -> 2
		Three = iota + 1 // iota 2 -> 3
		Four  = iota + 1 // iota 3 -> 4
	)

	output.Stdoutl("iota in the middle", One, Two, Three, Four)

	// iota を同じ行に複数書いても同じ値となる
	const (
		A, B, C = iota, iota, iota // 0, 0, 0
		D, E, F                    // 1, 1, 1
	)

	output.Stdoutl("iota multiple", A, B, C, D, E, F)

	// iota を同じ行に複数書くと同じ iota の値を受け取れる
	const (
		V1, V2 = iota, iota + 10 // 0, 10
		V3, V4                   // 1, 11
		// この場合、以下のように一行で一つだけの宣言は *そのまま* できない. コンパイルエラーとなる
		// 上で一行に2つの iota のパターンを利用しているため
		// V5,

		V5 = iota // iota をリセットするのであればオッケイ
	)

	output.Stdoutl("iota multiple2", V1, V2, V3, V4, V5)

	// Go の const では 直近の表現がキープされる
	// なので、以下のように 直値 で 1 と指定して次の行に値を指定しない場合
	// 直近の表現、つまり 1 がそのままキープされる。
	//
	// なので、iota を指定している場合は、暗黙で次の行でも iota が指定されていることになる
	const (
		First  = 1
		Second // Second = 1 と同じ
	)

	output.Stdoutl("last used expression keep", First, Second)

	// iota で フラグ向けの値はよく利用されるパターン
	// ビット操作で利用することが多い
	type (
		Flg int
	)
	const (
		Flg1    Flg = 1 << iota // 1 << 0 --> 0b00000001 --> 1
		Flg2                    // 1 << 1 --> 0b00000010 --> 2
		Flg3                    // 1 << 2 --> 0b00000100 --> 4
		AllFlgs = Flg1 | Flg2 | Flg3
	)

	flg := Flg2 | Flg3
	output.Stdoutl("iota bitwise value is", flg)

	if flg&Flg1 == Flg1 {
		output.Stdoutl("flg & Flg1", "Flg1")
	}

	if flg&Flg2 == Flg2 {
		output.Stdoutl("flg & Flg2", "Flg2")
	}

	if flg&Flg3 == Flg3 {
		output.Stdoutl("flg & Flg3", "Flg3")
	}

	if flg&AllFlgs != 0 {
		output.Stdoutl("flg & AllFlgs", "flg is valid")
	}

	wrongFlg := Flg(8) // 8 --> 1 << 3 --> 0b00001000
	if wrongFlg&AllFlgs == 0 {
		output.Stdoutl("Flg(8) & AllFlgs", "Flg(8) is invalid")
	}

	return nil
}
