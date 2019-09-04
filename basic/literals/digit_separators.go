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
}
