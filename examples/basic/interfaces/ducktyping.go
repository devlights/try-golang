package interfaces

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

type myType string

func (me myType) Say() string {
	return fmt.Sprintf("%T", me)
}

// DuckTyping2 -- ダックタイピングのちょっとしたサンプルです。
//
// REFERENCES:
//   - https://thomasnguyen.hashnode.dev/duck-typing-in-go
func DuckTyping2() error {
	var (
		fn = func(obj interface{ Say() string }) {
			output.Stdoutl("[fn]", obj.Say())
		}
	)

	// If it looks like a duck, and it quacks like a duck, then it is a duck.
	// アヒルのように見えて、アヒルのように鳴いていたら、それはアヒルである.
	fn(myType(""))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: interface_ducktyping2

	   [Name] "interface_ducktyping2"
	   [fn]                 interfaces.myType


	   [Elapsed] 9.56µs
	*/

}
