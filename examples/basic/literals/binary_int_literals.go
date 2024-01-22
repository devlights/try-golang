package literals

import "fmt"

// BinaryIntLiterals は、 Go1.13 から追加された Binary Integer Literals のサンプルです.
func BinaryIntLiterals() error {
	var (
		// 0b で始まると 2進数 と認識される
		bil1 = 0b1011
		// 0B で始めても同様
		bil2 = 0b1011
		// _ で区切ることもできる
		bil3 = 0b_10_11
		bil4 = 0b_10_11
	)

	fmt.Printf("%v\t%v\t%v\t%v\n", bil1, bil2, bil3, bil4)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: literals_binary_int_literals

	   [Name] "binary_int_literals"
	   11      11      11      11


	   [Elapsed] 19.03µs
	*/

}
