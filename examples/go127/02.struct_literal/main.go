// 構造体の埋め込みフィールドの簡易化 (構造体リテラルのキー拡張) のサンプルです。
//
// # REFERENCES
//   - https://victoriametrics.com/blog/go-1-27/#struct-literal-field-selectors
//   - https://future-architect.github.io/articles/20260805a/
package main

import (
	"fmt"
	"uuid"
)

type (
	Base struct {
		ID uuid.UUID
	}

	Item[T any] struct {
		Base
		Value T
	}
)

func (i Item[T]) String() string {
	return fmt.Sprintf("Item[ID:%s, Value:%v]", i.ID, i.Value)
}

func main() {
	// Go 1.26まで
	i1 := Item[int]{Base: Base{ID: uuid.NewV7()}, Value: 42}
	// Go 1.27から
	i2 := Item[int]{ID: uuid.NewV7(), Value: 42}

	fmt.Println(i1, i2)
}
