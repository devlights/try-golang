package reflection

import (
	"fmt"
	"reflect"
)

// TypeOf -- reflect.TypeOf() のサンプル
//noinspection GoNameStartsWithPackageName
func TypeOf() error {
	// 特定の値の型を調べる場合には、reflect.TypeOf() を利用する
	// C# でいう、 "hello world".GetType() と同じような感じ
	// Name メソッドで名前を取得できる
	i := 0
	t1 := reflect.TypeOf(i)
	fmt.Printf("%s\n", t1.Name())

	s := "hello world"
	t2 := reflect.TypeOf(s)
	fmt.Printf("%s\n", t2.Name())

	return nil
}