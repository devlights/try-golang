package mapop

import "fmt"

// MapInitialize -- マップの初期化に関するサンプルです。
func MapInitialize() error {

	// var で 定義して make で割当
	var m1 map[string]int //lint:ignore S1021 サンプルなので意図的にこのようにしている
	m1 = make(map[string]int)

	fmt.Printf("[m1] len: %d\tval: %#v\n", len(m1), m1)

	m1["a"] = 100
	m1["b"] = 200

	fmt.Printf("[m1] len: %d\tval: %v\n", len(m1), m1)

	// 直接初期値を設定して作る
	m2 := map[string]int{
		"a": 100,
		"b": 200,
	}

	fmt.Printf("[m2] len: %d\tval: %v\n", len(m2), m2)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: map_initialize

	   [Name] "map_initialize"
	   [m1] len: 0     val: map[string]int{}
	   [m1] len: 2     val: map[a:100 b:200]
	   [m2] len: 2     val: map[a:100 b:200]


	   [Elapsed] 45µs
	*/

}
