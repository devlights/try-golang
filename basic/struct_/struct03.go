package struct_

import "fmt"

type Base struct {
	id   int
	name string
}

type A struct {
	Base   // 共通フィールド
	ValueA string
}

type B struct {
	Base   // 共通フィールド
	ValueB string
}

// 構造体間で共有するフィールド郡を共通化
func Struct03() error {
	// 複合リテラルを用いて初期化する場合
	// 以下のようにBaseの部分は明示的に
	// Base分として指定が必要
	a := A{
		Base: Base{
			id:   100,
			name: "A",
		},
		ValueA: "val-a",
	}

	// 以下のようにひとつずつフィールドを
	// 埋めていく場合、Baseの部分は省略して
	// 記載することが可能
	var b B
	b.id = 200   // b.Base.id と同じ
	b.name = "B" // b.Base.name と同じ
	b.ValueB = "val-b"

	fmt.Println(a, b)

	return nil
}
