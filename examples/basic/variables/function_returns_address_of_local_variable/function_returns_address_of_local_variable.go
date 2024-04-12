package function_returns_address_of_local_variable

import (
	"github.com/devlights/gomy/output"
)

// FunctionReturnsAddressOfLocalVariable -- Goでは関数がローカル変数のアドレスを返すのは全く問題ないことを示すサンプルです.
func FunctionReturnsAddressOfLocalVariable() error {
	// -----------------------------------------------
	// Go では、関数がローカル変数のアドレスを返すのは全く問題ない.
	// (書籍「プログラミング言語 Go」初版 P.35より)
	//
	// C  では、関数のローカル変数は関数からリターンした時点でメモリから開放されてしまうので
	// 呼び出し元が戻り値として受け取ったアドレスに大してデリファレンスすると
	// 高確率で segmentation fault する.
	//
	// Go はGCを備えているので、気にせず関数から生成したオブジェクトのポインタを
	// 返しても、何の問題もない。
	// -----------------------------------------------
	i1 := f()
	output.Stdoutf("[i1]", "%p\t%v\n", i1, *i1)

	i2 := f()
	output.Stdoutf("[i2]", "%p\t%v\n", i2, *i2)

	// i1, i2 ともに関数fの内部でローカル変数として生成されたものであるが、ちゃんと存在している
	// C 言語の場合は、呼び出し元に返ってきた時点でアドレスが示す先のメモリ領域は開放されているため
	// 値は不定となる。

	f2(i1, i2)
	output.Stdoutf("[i1/i2]", "%v\t%v\n", *i1, *i2)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: function_returns_address_of_local_variable

	   [Name] "function_returns_address_of_local_variable"
	   [i1]                 0xc0001a2908       1
	   [i2]                 0xc0001a2920       1
	   [i1/i2]              2  2


	   [Elapsed] 32.3µs
	*/

}

func f() *int {
	i := 1
	return &i
}

func f2(i1, i2 *int) {
	*i1++
	*i2++
}
