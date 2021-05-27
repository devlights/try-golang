package maps

import "fmt"

// MapDelete -- マップの要素を削除するサンプルです。
func MapDelete() error {

	var (
		f = func(v map[string]int) {
			fmt.Printf("len: %d\tval: %v\n", len(v), v)
		}

		m = map[string]int{
			"a": 100,
			"b": 200,
		}
	)

	f(m)

	// map から要素を削除する場合 組み込み関数 delete() を使う
	delete(m, "a")

	f(m)

	delete(m, "b")

	f(m)

	// 存在しない要素を delete に渡してもエラーにはならない。何も起きないだけ。
	delete(m, "a")

	f(m)

	return nil
}
