package literals

import "fmt"

// HexIntLiterals は、 go の 16 進数リテラル のサンプルです.
func HexIntLiterals() error {
	var (
		// 0x で始まると 8進数 と認識される
		hex1 = 0xff
		// 0X で始めても同様
		hex2 = 0xff
		// _ で区切ることもできる
		hex3 = 0x_f_f
		hex4 = 0x_f_f
	)

	fmt.Printf("%v\t%v\t%v\t%v\n\n", hex1, hex2, hex3, hex4)

	return nil
}
