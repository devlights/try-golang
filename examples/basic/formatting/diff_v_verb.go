package formatting

import "github.com/devlights/gomy/output"

// DiffVVerb は、v verbの違いについてのサンプルです.
func DiffVVerb() error {
	type F struct {
		id   int
		name string
	}

	f := &F{100, "helloworld"}

	output.Stdoutf("[%v ]", "%v \n", f) // 型名無し、フィールド名無し
	output.Stdoutf("[%+v]", "%+v\n", f) // 型名無し、フィールド名有り
	output.Stdoutf("[%#v]", "%#v\n", f) // 型名有り、フィールド名有り
	output.Stdoutf("[%T ]", "%T \n", f)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: formatting_diff_v_verb

	   [Name] "formatting_diff_v_verb"
	   [%v ]                &{100 helloworld}
	   [%+v]                &{id:100 name:helloworld}
	   [%#v]                &formatting.F{id:100, name:"helloworld"}
	   [%T ]                *formatting.F


	   [Elapsed] 41.59µs
	*/

}
