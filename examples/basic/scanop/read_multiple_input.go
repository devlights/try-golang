package scanop

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

// ReadMultipleInput は、fmt.Scan() で複数の値を読み取るサンプルです.
//
// # REFERENCES
//
//   - https://dev.to/azure/go-from-the-beginning-reading-user-input-i79
//   - https://pkg.go.dev/fmt@go1.19.2#Scan
func ReadMultipleInput() error {
	var (
		value1, value2 string
	)

	fmt.Print("INPUT: ")

	// 改行は一つの空白（つまり区切り文字）として認識される
	// つまり "hello world" と
	// "hello
	//
	// world"
	// は同じ挙動となる.
	n, err := fmt.Scan(&value1, &value2)
	if err != nil {
		return err
	}

	output.Stdoutf("[fmt.Scan]", "count=%d\tvalue1=%v\tvalue2=%v\n", n, value1, value2)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: scanop_read_multi_input

	   [Name] "scanop_read_multi_input"
	   INPUT: helloworld
	   helloworld2
	   [fmt.Scan]           count=2    value1=helloworld       value2=helloworld2


	   [Elapsed] 6.552806734s
	*/

}
