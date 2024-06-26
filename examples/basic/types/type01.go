package types

import "fmt"

// MyInt1 -- int に対して別名を定義
type MyInt1 int

// MyInt2 -- int に対して別名を定義
type MyInt2 int

// MyIntArray -- 配列に対して別名を定義
type MyIntArray [3]int

// KeyValuePair -- 構造体の定義も struct で定義した内容を type で名前定義しているのと同じ
type KeyValuePair struct {
	Key   string
	Value interface{}
}

// NewKeyValuePair -- *KeyValuePair を生成
func NewKeyValuePair(k string, v string) *KeyValuePair {
	return &KeyValuePair{Key: k, Value: v}
}

// Basic は、type についてのサンプルです.
func Basic() error {
	i1 := MyInt1(100)
	i2 := MyInt2(200)
	a1 := MyIntArray{1, 2, 3}
	s1 := NewKeyValuePair("hello", "world")

	fmt.Println(i1, i2, a1, s1)
	fmt.Printf("%T, %T, %T, %T\n", i1, i2, a1, s1)

	// Goでは型が厳密にチェックされる
	// MyInt1はintの別名定義なので int に変換可能
	i3 := int(i1)
	fmt.Println(i1, i3)

	// 同じ int の別名定義であっても MyInt1 と MyInt2 は全く別の型と認識される
	// つまり、同じ型から派生した場合でも、エイリアス間にはには互換性がない
	// 以下はコンパイルエラー
	// i2 = i1
	// 以下は通る
	i2 = MyInt2(i1)
	fmt.Println(i2)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: types_basic

	   [Name] "types_basic"
	   100 200 [1 2 3] &{hello world}
	   types.MyInt1, types.MyInt2, types.MyIntArray, *types.KeyValuePair
	   100 100
	   100


	   [Elapsed] 35.82µs
	*/

}
