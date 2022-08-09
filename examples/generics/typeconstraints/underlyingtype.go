package typeconstraints

import "github.com/devlights/gomy/output"

type signed interface {
	~int8 | ~int16 | ~int32 | ~int64
}

type unsigned interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

type my8 int8
type my32 uint32

func add1[T signed](x, y T) T {
	return x + y
}

func add2[T unsigned](x, y T) T {
	return x + y
}

// UnderlyingType -- Generics時に利用できる ~トークン についてのサンプルです。
//
// REFERENCES
//   - https://go.dev/tour/generics/1
//   - https://go.dev/tour/generics/2
//   - https://go.dev/doc/tutorial/generics
//   - https://go.dev/blog/intro-generics
//   - https://go.dev/blog/when-generics
func UnderlyingType() error {
	output.Stdoutl("[int8, int8]", add1(int8(1), int8(5)))
	output.Stdoutl("[my8, my8]", add1(my8(1), my8(5)))

	// コンパイルエラー
	//   int32 does not implement unsigned
	//output.Stdoutl("[int32, int32]", add2(int32(1), int32(5)))

	output.Stdoutl("[int32, int32]", add2(uint32(1), uint32(5)))
	output.Stdoutl("[my32, my32]", add2(my32(1), my32(5)))

	return nil
}
