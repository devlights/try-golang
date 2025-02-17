package go124

import "fmt"

type (
	Item[T any] struct {
		Value T
	}

	// Generic Type Definition (完全に新しい型を作成) -- これは Go 1.24 以前から出来ていた
	TypeDefinition[T any] Item[T]
	// Generic Type Alias      (既存の型に別名を付与) -- これが Go 1.24 から出来るようになった
	TypeAlias[T any] = Item[T]
)

func newTypeDefinition[T any](v T) *TypeDefinition[T] {
	return &TypeDefinition[T]{Value: v}
}

func newTypeAlias[T any](v T) *TypeAlias[T] {
	return &TypeAlias[T]{Value: v}
}

// Type Definition は、完全に新しい型を作成するので、独自メソッドを追加可能
func (me *TypeDefinition[T]) String() string {
	return fmt.Sprintf("(%v:%p)", me.Value, &me.Value)
}

// Type Alias は、元の型の別名なので、独自メソッドの追加不可能
// (cannot define new methods on generic alias type TypeAlias[T any]compiler (InvalidRecv))
//
// func (me *TypeAlias[T]) String() string {
// 	return fmt.Sprintf("(%v:%p)", me.Value, &me)
// }

// GenericTypeAlias は、Go 1.24 で追加された Generic Type Alias のサンプルです.
//
// GenericなType Definitionは、Go 1.24以前でも可能でしたが
// Go 1.24にて、GenericなType Aliasも可能となりました。([Go 1.24 Release note])
//
//   - Type Definition は、完全に新たな型を作成する機能です。元の型との互換性はありません。
//   - Type Alias      は、元の型の別名を作成する機能です。元の型と互換性があります。
//
// コードでは、以下のようになります。
//
//	type ID int   // Type Definition
//	type ID = int // Type Alias
//
// [Go 1.24 Release note]: https://tip.golang.org/doc/go1.24
func GenericTypeAlias() error {
	var (
		td = newTypeDefinition(100)
		ta = newTypeAlias(200)

		item *Item[int]
	)

	// Type Definitionは元の型との互換性が無いためコンパイルエラーとなる
	// (cannot use td (variable of type *TypeDefinition[int]) as *Item[int] value in assignment compiler (IncompatibleAssign))
	//
	//item = td

	// Type Aliasは、元の型の別名であるため互換性がある
	item = ta

	fmt.Printf("%s\t(%v:%p)\t(%v:%p)\n", td, ta.Value, &ta.Value, item.Value, &item.Value)

	return nil

	/*
	   $ task
	   task: [build] go build -o "/workspace/try-golang/try-golang" .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: go124_generic_type_alias

	   [Name] "go124_generic_type_alias"
	   (100:0xc000012a70)      (200:0xc000012a78)      (200:0xc000012a78)

	   [Elapsed] 19.6µs
	*/
}
