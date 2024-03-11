package structs

import (
	"fmt"
	"reflect"
)

//lint:ignore U1000 It's ok because this is just a example.
type intPair struct {
	x  int
	y  int
	s  string
	sl []int
}

// Basic02 -- 構造体が値型であることの確認
// noinspection GoNilness
func Basic02() error {

	// 構造体は値型なので宣言した時点でメモリ上に領域が確保される
	// 各フィールドは初期値で初期化される
	var s01 intPair
	fmt.Printf("%#v\n", s01)

	// スライスは参照なので宣言した時点では nil
	var s02 []int
	fmt.Printf("%#v is nil? ==> %v\n", s02, s02 == nil)

	// 参照型の場合は、reflect.ValueOf(x).IsNil() でも判定できる
	// ただし、参照型以外を渡すと panic になる
	isNil := reflect.ValueOf(s02).IsNil()
	fmt.Printf("%#v\n", isNil)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: struct_basic02

	   [Name] "struct_basic02"
	   structs.intPair{x:0, y:0, s:"", sl:[]int(nil)}
	   []int(nil) is nil? ==> true
	   true


	   [Elapsed] 32.46µs
	*/

}
