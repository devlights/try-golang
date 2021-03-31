// staticcheck の 警告 を無視するやり方についてのサンプルです.
//
// REFERENCES:
//   - https://staticcheck.io/docs#line-based-linter-directives
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./go.mod")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	//lint:ignore SA4006 Acknowledged. Okey.
	err = fn1()
}

func fn1() error {
	return nil
}
