package main

import "log"

func toptr[T any](v T) *T {
	return &v
}

func main() {
	//
	// Go 1.26 にて go fix が go vet と同じ Go Analysis Framework(golang.org/x/tools/go/analysis) を
	// 利用するようになった。以前までは go fix は独自の仕組みで動いていましたが、今後は go vet と同じ基盤を
	// 使って処理するようになり、モダンな修正提案・自動適用が出来るようになっている。
	//
	// Go 1.26 で追加された newexpr ルール (call of toptr(x) can be simplified to new(x)) も適用となるため
	// 以前までよく実装されていたポインタを取るだけの関数も go fix すると newexpr に自動的に置き換えられる。
	//

	log.SetFlags(0)
	log.SetPrefix("[main] ")

	var (
		v1 int32  = 1
		v2 string = "hello"

		p1 = toptr(v1)
		p2 = toptr(v2)
	)
	log.Printf("v=%[1]v(%[1]T), p=%[2]p(%[2]T)", v1, p1)
	log.Printf("v=%[1]v(%[1]T), p=%[2]p(%[2]T)", v2, p2)
}
