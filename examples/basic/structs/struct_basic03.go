package structs

import "fmt"

type baseType struct {
	id   int
	name string
}

type stA struct {
	baseType // 共通フィールド
	ValueA   string
}

type stB struct {
	baseType // 共通フィールド
	ValueB   string
}

// Basic03 -- 構造体間で共有するフィールド郡を共通化
func Basic03() error {
	// 複合リテラルを用いて初期化する場合
	// 以下のようにBaseの部分は明示的に
	// Base分として指定が必要
	a := stA{
		baseType: baseType{
			id:   100,
			name: "A",
		},
		ValueA: "val-a",
	}

	// 以下のようにひとつずつフィールドを
	// 埋めていく場合、Baseの部分は省略して
	// 記載することが可能
	var b stB
	b.id = 200   // b.Base.id と同じ
	b.name = "B" // b.Base.name と同じ
	b.ValueB = "val-b"

	fmt.Println(a, b)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: struct_basic03

	   [Name] "struct_basic03"
	   {{100 A} val-a} {{200 B} val-b}


	   [Elapsed] 12.12µs
	*/

}
