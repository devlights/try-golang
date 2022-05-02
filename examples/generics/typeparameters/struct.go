package typeparameters

import "github.com/devlights/gomy/output"

type genericType[T any] struct {
	value T
}

// Struct -- Type Parameterを構造体に適用するサンプルです.
//
// メモ
//   無名構造体はジェネリックに出来ない。コンパイルエラーとなる。
//   https://groups.google.com/g/golang-nuts/c/ulntjLeGYn4/m/e3dWOOqtAwAJ
//
//   VSCode で golang.go 拡張機能を利用している場合、カーソルを載せた際に表示される
//   ツールチップにはジェネリックの場合、具象化されたシグネチャが表示されるので便利。
//
// REFERENCES
//  - https://go.dev/tour/generics/1
//  - https://go.dev/tour/generics/2
//  - https://go.dev/tour/generics/3
//  - https://go.dev/doc/tutorial/generics
//  - https://go.dev/blog/intro-generics
//  - https://go.dev/blog/when-generics
func Struct() error {
	var (
		s1 = genericType[int]{100}
		s2 = genericType[string]{"helloworld"}
		s3 = genericType[genericType[int]]{s1}
	)

	output.Stdoutf("[s1]", "%+v\t(%T)\n", s1, s1)
	output.Stdoutf("[s2]", "%+v\t(%T)\n", s2, s2)
	output.Stdoutf("[s3]", "%+v\t(%T)\n", s3, s3)

	return nil
}
