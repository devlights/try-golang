package structs

import "fmt"

type mySt01 struct {
	key   int
	value string
}

// Basic04 -- 組み込み関数 new() のサンプル
func Basic04() error {

	st01 := &mySt01{
		key:   100,
		value: "hello world",
	}

	// 組み込み関数 new() は指定した型のポインタを生成する関数
	st02 := new(mySt01)
	st02.key = 200
	st02.value = "world hello"

	// 当然既存の値型にも利用できる
	i01 := new(int)
	*i01 = 111

	fmt.Printf("%#v\n%#v\n%#v(%#v)\n", st01, st02, i01, *i01)

	return nil
}
