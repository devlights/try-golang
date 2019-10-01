package tutorial

import "fmt"

// Struct は、 Tour of Go - Structs (https://tour.golang.org/moretypes/2) の サンプルです。
func Struct() error {
	// ------------------------------------------------------------
	// Go言語の構造体
	//
	// 構造体は、フィールドの集まりを表現する.
	// Go言語にて、構造体は type キーワード を利用して定義する.
	//
	// ちなみに、Go言語において struct{} は 「空構造体」 を表す
	// 空構造体はメモリを消費しない.
	// ------------------------------------------------------------
	type (
		point struct {
			x, y int
		}
	)

	//noinspection GoVarAndConstTypeMayBeOmitted
	var (
		// 構造体は、基本ポインタで扱うのが多い
		p1 = &point{x: 10, y: 20}
		p2 point  // 初期値を設定していない場合は、値はゼロ値となる
		p3 *point // ポインタの初期値は nil
	)

	// 文字列書式に %v をつけると 値 が表示される
	// %#v をつけると、値と型も表示される
	fmt.Printf("%#v\n", p1)

	// p1 は、ポインタであるため、普通は デリファレンス が必要となる
	// (*p1).x と記載しないといけないはず。
	// だが、この表記法は面倒なので、言語側で特殊措置をしてくれている
	// そのため、普通の値の場合と同様に, p1.x と書いてもちゃんと動く.
	fmt.Printf("x:%v\ty:%v\n", p1.x, p1.y)

	// 配列の実体は nil にはならない。（初期値は Zero Value となる)
	// ポインタの初期値は nil
	fmt.Printf("p2 is nil? => %v\tp3 is nil? => %v\n", p2, p3)

	return nil
}
