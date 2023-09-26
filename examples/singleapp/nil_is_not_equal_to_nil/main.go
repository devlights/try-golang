package main

import (
	"github.com/devlights/gomy/output"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	//
	// Go では、nil 同士を比較しても同じにならない場合がある
	//
	// REFERENCES:
	//	- https://www.calhoun.io/when-nil-isnt-equal-to-nil/
	//  - https://forum.golangbridge.org/t/a-nil-a-b-b-nil-with-pointers-and-interface/10593
	//  - https://staticcheck.io/docs/checks#SA4023
	//  - https://devlights.hatenablog.com/entry/2021/03/31/235948
	//

	var (
		v1 *int // ゼロ値は nil
		v2 any  // ゼロ値は nil
	)

	// どちらもnil
	output.Stdoutl("[v1 is nil?]", v1 == nil)
	output.Stdoutl("[v2 is nil?]", v2 == nil)

	// 同じ？ (結果は false となる)
	//
	// staticcheck では以下の警告が出る
	//   this comparison is never true; the lhs of the comparison has been assigned a concretely typed value (SA4023)
	//
	output.Stdoutl("[v1 eq v2]", v1 == v2) //lint:ignore SA4023 It's ok because this is just a example.

	output.StdoutHr()

	//
	// v1をv2に設定してみる
	// (論理的には nil な変数に nil を設定していることになる)
	//
	v2 = v1

	// 同じ？ (結果は true となる)
	//
	// staticcheckでは今度は v2 == nil の部分で警告が出る
	//   this comparison is never true; the lhs of the comparison has been assigned a concretely typed value (SA4023)
	//
	output.Stdoutl("[v2=v1][v1 is nil?]", v1 == nil)
	output.Stdoutl("[v2=v1][v2 is nil?]", v2 == nil) //lint:ignore SA4023 It's ok because this is just a example.
	output.Stdoutl("[v2=v1][v1 eq v2]", v1 == v2)

	output.StdoutHr()

	// 大事な点として、Goのすべてのポインタには2つの基本情報があるということ。
	// ポインターの型と、ポインターが指す値である。
	var (
		v3 *int
		v4 any
	)

	output.Stdoutf("[v3]", "(%T,%v)\n", v3, v3) // (*int, <nil>)
	output.Stdoutf("[v4]", "(%T,%v)\n", v4, v4) // (<nil>, <nil>)

	// v3とv4を (型, 値) として見ると同じではない
	// なので、同じかどうかを聞くと同じとならない
	output.Stdoutl("[v3 eq v4]", v3 == v4) //lint:ignore SA4023 It's ok because this is just a example.
	output.StdoutHr()

	// ここで v3 を v4 に設定すると、値はnilのままだが、型が設定されるため同じになる
	v4 = v3
	output.Stdoutf("[v3]", "(%T,%v)\n", v3, v3) // (*int, <nil>)
	output.Stdoutf("[v4]", "(%T,%v)\n", v4, v4) // (*int, <nil>)
	output.Stdoutl("[v3 eq v4]", v3 == v4)
	output.StdoutHr()

	// v4 は any, つまり interface{} である
	// Go のインターフェースは 型と値 が、両方 nil の場合に nil となる
	// つまり型だけが設定されている場合は nil とならない
	output.Stdoutl("[v4 == nil?]", v4 == nil) //lint:ignore SA4023 It's ok because this is just a example.

	//
	// イディオムとして覚えておくことは、nilを設定したい場合は
	// 「nilな変数」を設定するのではなく、明示的に 「nil」 を設定すること
	//
	v4 = nil
	output.Stdoutf("[v4 = nil][v3]", "(%T,%v)\n", v3, v3) // (*int, <nil>)
	output.Stdoutf("[v4 = nil][v4]", "(%T,%v)\n", v4, v4) // (<nil>, <nil>)
	output.Stdoutl("[v4 = nil][v3 eq v4]", v3 == v4)      //lint:ignore SA4023 It's ok because this is just a example.

	// --------------------------------------------------
	// 実行結果
	// --------------------------------------------------
	// [v1 is nil?]         true
	// [v2 is nil?]         true
	// [v1 eq v2]           false
	// --------------------------------------------------
	// [v2=v1][v1 is nil?]  true
	// [v2=v1][v2 is nil?]  false
	// [v2=v1][v1 eq v2]    true
	// --------------------------------------------------
	// [v3]                 (*int,<nil>)
	// [v4]                 (<nil>,<nil>)
	// [v3 eq v4]           false
	// --------------------------------------------------
	// [v3]                 (*int,<nil>)
	// [v4]                 (*int,<nil>)
	// [v3 eq v4]           true
	// --------------------------------------------------
	// [v4 == nil?]         false
	// [v4 = nil][v3]       (*int,<nil>)
	// [v4 = nil][v4]       (<nil>,<nil>)
	// [v4 = nil][v3 eq v4] false

	return nil
}
