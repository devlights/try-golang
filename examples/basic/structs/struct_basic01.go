package structs

import "fmt"

// MyStruct -- サンプル用の構造体
type MyStruct struct {
	// 値
	Value int
}

// 構造体にメソッドを追加する場合は外側で関数を定義して紐付ける
// 紐付ける際に設定するレシーバーの型により、紐づけ先が異なる

// Method1 -- レシーバーがポインタの場合のメソッド
func (m *MyStruct) Method1() int {
	return m.Value
}

// Method2 -- レシーバーがポインタではない場合のメソッド
func (m MyStruct) Method2() int {
	return m.Value * 2
}

func updateValue(m MyStruct) {
	m.Value += 10
}

func updateValuePtr(m *MyStruct) {
	m.Value += 10
}

// Basic01 -- 構造体についてのサンプル
//noinspection GoUnhandledErrorResult
func Basic01() error {
	// ポインタと通常の2パターンを生成
	st01 := &MyStruct{
		Value: 30,
	}

	st02 := MyStruct{
		Value: 55,
	}

	// 値を表示
	pf := fmt.Printf
	pf("Method1: %d(%T)\n", st01.Method1(), st01) // -> 30
	pf("Method2: %d(%T)\n", st01.Method2(), st01) // -> 60
	pf("Method1: %d(%T)\n", st02.Method1(), st02) // -> 55
	pf("Method2: %d(%T)\n", st02.Method2(), st02) // -> 110

	// それぞれの値を更新
	updateValuePtr(st01)
	updateValue(st02)
	fmt.Printf("\n\n")

	// 再度値を表示
	// 当然ながら、ポインタでは無い方の構造体は値が変化しない
	pf("Method1: %d(%T)\n", st01.Method1(), st01) // -> 40
	pf("Method2: %d(%T)\n", st01.Method2(), st01) // -> 80
	pf("Method1: %d(%T)\n", st02.Method1(), st02) // -> 55
	pf("Method2: %d(%T)\n", st02.Method2(), st02) // -> 110

	return nil
}
