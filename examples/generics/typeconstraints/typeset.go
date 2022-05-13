package typeconstraints

import "github.com/devlights/gomy/output"

type num interface {
	int32 | int64
}

func f1[T num](x, y T) T {
	return x + y
}

// TypeSet -- Generics 時に利用できる 型セット についてのサンプルです。
//
// REFERENCES
//  - https://go.dev/tour/generics/1
//  - https://go.dev/tour/generics/2
//  - https://go.dev/doc/tutorial/generics
//  - https://go.dev/blog/intro-generics
//  - https://go.dev/blog/when-generics
func TypeSet() error {
	output.Stdoutl("[f1]", f1(int32(10), int32(20)))
	output.Stdoutl("[f1]", f1(int64(10), int64(20)))

	return nil
}
