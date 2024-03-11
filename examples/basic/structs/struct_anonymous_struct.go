package structs

import "fmt"

// StructAnonymousStruct -- 匿名構造体についてのサンプルです。
func StructAnonymousStruct() error {
	// ------------------------------------------------------------
	// Go言語では、匿名なstructを構築することが出来る
	// 要は宣言と初期化を一気に済ませる
	// References:: http://bit.ly/2Lr0oq9
	// ------------------------------------------------------------
	anonStruct := struct {
		x, y int
	}{
		x: 100,
		y: 200,
	}

	fmt.Printf("[anonStruct] %#v\n", anonStruct)

	anonStructSlice := []struct {
		x, y int
	}{
		{x: 100, y: 200},
		{x: 300, y: 400},
	}

	fmt.Printf("[anonStructSlice] %#v\n", anonStructSlice)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: struct_anonymous_struct

	   [Name] "struct_anonymous_struct"
	   [anonStruct] struct { x int; y int }{x:100, y:200}
	   [anonStructSlice] []struct { x int; y int }{struct { x int; y int }{x:100, y:200}, struct { x int; y int }{x:300, y:400}}


	   [Elapsed] 21.84µs
	*/

}
