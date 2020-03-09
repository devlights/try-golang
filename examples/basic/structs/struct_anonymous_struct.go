package structs

import "fmt"

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
}
