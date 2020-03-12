package effectivego

import (
	"fmt"
	"strings"
)

// Functions -- Effective Go - Functions の 内容についてのサンプルです。
func Functions() error {
	/*
		https://golang.org/doc/effective_go.html#functions

		- Goの関数は複数の戻り値を返すことができる
		- Goの関数の戻り値には名前を予め付与することができる
	*/

	// 複数戻り値
	x, y := 10, 20
	x2, y2, err := multipleReturnValues(x, y)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(x2, y2)

	// 名前付き戻り値
	s := namedReturnValue("hello")
	fmt.Println(s)

	return nil
}

func multipleReturnValues(x int, y int) (int, int, error) {
	var err error
	if x < 0 || y < 0 {
		err = fmt.Errorf("negative x[%d] y[%d]", x, y)
	}

	return x * 2, y * 2, err
}

func namedReturnValue(s string) (r string) {
	r = strings.ToUpper(s)

	// 戻り値に名前を付与している場合、returnには何も付けなくて良い
	// return したタイミングの変数の値が戻り値として返る
	return
}
