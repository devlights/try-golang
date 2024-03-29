package stdout

import "fmt"

type myType struct {
	i int
	v int
	x int
	y int
}

// Printf02 -- 標準出力についてのサンプル
func Printf02() error {
	// ------------------------------------------------------
	// fmt.Printf() は　C言語の stdout 関数と
	// 同じような使い勝手を提供する。 書式については
	// 以下を参照。
	//     https://golang.org/fmt/#hdr-Printing
	// ------------------------------------------------------
	// %s は文字列
	messageFormat := "[%%s] Hello %s\n"
	fmt.Printf(messageFormat, "World")

	// %d は数値
	messageFormat = "[%%d] Hello %d\n"
	fmt.Printf(messageFormat, 100)

	// %v は デフォルトのフォーマットを用いて値を表示
	messageFormat = "[%%v] Hello %v\n"
	fmt.Printf(messageFormat, myType{i: 100, v: 111, x: 222, y: 333})

	// %#v は、Goのリテラル表現で値を出力
	messageFormat = "[%%#v] Hello %#v\n"
	fmt.Printf(messageFormat, myType{i: 100, v: 111, x: 222, y: 333})

	// %+v は %v に加えて構造体の場合にフィールド名も出力
	messageFormat = "[%%+v] Hello %+v\n"
	fmt.Printf(messageFormat, myType{i: 100, v: 111, x: 222, y: 333})

	// %T は データの型情報を出力
	data := myType{i: 100, v: 111, x: 222, y: 333}
	intArray := [...]int{1, 2, 3}
	intSlice := []int{1, 2, 3}
	intKeyMap := map[int]string{1: "apple"}
	messageFormat = "[%%T] Hello %T, %T, %T, %T, %T, %T\n"
	fmt.Printf(messageFormat, data, "hello", 100, intArray, intSlice, intKeyMap)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: printf02

	   [Name] "printf02"
	   [%s] Hello World
	   [%d] Hello 100
	   [%v] Hello {100 111 222 333}
	   [%#v] Hello stdout.myType{i:100, v:111, x:222, y:333}
	   [%+v] Hello {i:100 v:111 x:222 y:333}
	   [%T] Hello stdout.myType, string, int, [3]int, []int, map[int]string


	   [Elapsed] 22.46µs
	*/

}
