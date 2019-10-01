package tutorial

import (
	"fmt"
)

// HelloWorld は、 [A Tour of Go](http://bit.ly/2HsCMiG) の 要約.
func HelloWorld() error {
	// ------------------------------------------------------------
	// Hello World
	//   文字列を出力するには、 fmt パッケージの Println() などを使う
	// ------------------------------------------------------------
	fmt.Println("Hello World")

	return nil
}
