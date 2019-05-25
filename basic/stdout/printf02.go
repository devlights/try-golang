package stdout

import "fmt"

type MyType struct {
	i int
	v int
	x int
	y int
}

func Printf02() error {
	// ------------------------------------------------------
	// fmt.Printf() は　C言語の stdout 関数と
	// 同じような使い勝手を提供する。 書式については
	// 以下を参照。
	//     https://golang.org/pkg/fmt/#hdr-Printing
	// ------------------------------------------------------
	// %s は文字列
	messageFormat := "Hello %s\n"
	fmt.Printf(messageFormat, "World")

	// %d は数値
	messageFormat = "Hello %d\n"
	fmt.Printf(messageFormat, 100)

	// %v は デフォルトのフォーマットを用いて値を表示
	messageFormat = "Hello %v\n"
	fmt.Printf(messageFormat, MyType{i: 100, v: 111, x: 222, y: 333})

	// %+v は %v に加えて構造体の場合にフィールド名も出力してくれる
	messageFormat = "Hello %+v\n"
	fmt.Printf(messageFormat, MyType{i: 100, v: 111, x: 222, y: 333})

	return nil
}
