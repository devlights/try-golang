package formatting

import "github.com/devlights/gomy/output"

// UsingV -- フォーマットするときの v の使い方についてのサンプルです。
//
// %v, %+v, %#v があります。
//
// REFERENCES
//   - https://www.developer.com/languages/structure-golang/
func UsingV() error {
	type S struct {
		Id   int
		Name string
	}

	v := S{100, "helloworld"}

	output.Stdoutf("[v ]", "%v\n", v)  // フィールドの値だけ
	output.Stdoutf("[+v]", "%+v\n", v) // フィールド名と値
	output.Stdoutf("[#v]", "%#v\n", v) // 型名とフィールド名と値

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: formatting_using_v

	   [Name] "formatting_using_v"
	   [v ]                 {100 helloworld}
	   [+v]                 {Id:100 Name:helloworld}
	   [#v]                 formatting.S{Id:100, Name:"helloworld"}


	   [Elapsed] 62.56µs
	*/

}
