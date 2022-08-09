package typeparameters

func equal1[T any](x T, y T) bool {
	// コンパイルエラー
	//   invalid operation: cannot compare x == y (type parameter T is not comparable with ==)
	//return x == y

	return true
}

func equal2[T comparable](x T, y T) bool {
	// コンパイルが通る
	// 型引数に comparable を指定しているので、比較可能であることを示している
	return x == y
}

// Comparable -- Go 1.18 より追加されたビルドイン制約型 comparable についてのサンプルです。
//
// REFERENCES
//   - https://go.dev/tour/generics/1
//   - https://go.dev/tour/generics/2
//   - https://go.dev/doc/tutorial/generics
//   - https://go.dev/blog/intro-generics
//   - https://go.dev/blog/when-generics
func Comparable() error {
	equal1("hello", "hello")
	equal2("world", "world")

	return nil
}
