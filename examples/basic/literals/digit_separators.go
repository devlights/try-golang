package literals

import "fmt"

// DigitSeparators は、 go1.13 の　Digit separators のサンプルです.
func DigitSeparators() error {
	var (
		i1  = 100_00
		i2  = 1_000_000
		i3  = 3.1415_9265
		arr = [3]interface{}{i1, i2, i3}
	)

	fmt.Printf("%v\n", arr)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: literals_digit_separator

	   [Name] "digit_separator"
	   [10000 1000000 3.14159265]


	   [Elapsed] 20.81µs
	*/

}
