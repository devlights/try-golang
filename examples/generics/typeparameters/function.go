package typeparameters

import "github.com/devlights/gomy/output"

// Function -- Type Parameterを関数に適用するサンプルです.
//
// メモ
//   無名関数はジェネリックに出来ない。コンパイルエラーとなる。
//   https://groups.google.com/g/golang-nuts/c/ulntjLeGYn4/m/e3dWOOqtAwAJ
//
//   VSCode で golang.go 拡張機能を利用している場合、関数にカーソルを載せた際に表示される
//   ツールチップにはジェネリック関数の場合、具象化されたシグネチャが表示されるので便利。
//
// REFERENCES
//  - https://go.dev/tour/generics/1
//  - https://go.dev/tour/generics/2
//  - https://go.dev/tour/generics/3
//  - https://go.dev/doc/tutorial/generics
//  - https://go.dev/blog/intro-generics
//  - https://go.dev/blog/when-generics
func Function() error {
	// Go 1.18 まではジェネリクスが存在しないので、型毎に関数を定義するか
	// インターフェースを使って抽象化する必要があった。(別段不便ではないが)

	var (
		i1 = int32(100)
		i2 = int64(110)
	)
	output.Stdoutl("[fn1 -- int32]", fn1(i1))
	output.Stdoutl("[fn2 -- int64]", fn2(i2))
	// 以下はコンパイルエラー
	//output.Stdoutl("[fn2]", fn2(i1))

	// Go 1.18 からはジェネリクスが搭載されたので以下のように書ける
	output.Stdoutl("[fng -- int32]", fng(i1))
	output.Stdoutl("[fng -- int64]", fng(i2))

	return nil
}

func fn1(i int32) int32 {
	return i
}

func fn2(i int64) int64 {
	return i
}

func fng[T int32 | int64](i T) T {
	return i
}
