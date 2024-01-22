package literals

import "fmt"

// OctalIntLiterals は、 Go1.13 から追加された Octal Integer Literals のサンプルです.
func OctalIntLiterals() error {
	var (
		// 0o で始まると 8進数 と認識される
		oil1 = 0o660
		// 0O で始めても同様
		oil2 = 0o660
		// _ で区切ることもできる
		oil3 = 0o_6_6_0
		oil4 = 0o_6_6_0
	)

	fmt.Printf("%v\t%v\t%v\t%v\n", oil1, oil2, oil3, oil4)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: literals_octal_int_literals

	   [Name] "octal_int_literals"
	   432     432     432     432


	   [Elapsed] 20.02µs
	*/

}
