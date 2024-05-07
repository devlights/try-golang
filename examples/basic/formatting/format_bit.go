package formatting

import (
	"fmt"
)

// Nbit は、書式表示の %b を利用して指定ビット数分を表示するサンプルです。
func Nbit() error {
	var (
		v1 = 127
		v2 = 1 << 8
		v3 = 1 << 10
		v4 = 0x5555
		v5 = 0xaaaa

		fn = func(v int) {
			// %0Nb で、指定ビット桁数で2進数表示できる
			// 以下は 16ビット 分で表示
			//
			// 当然、指定桁より大きい場合も表示されるが、表示がずれるので注意
			fmt.Printf("%016b\t0x%04[1]x\t%5[1]d\n", v)
		}
	)

	for _, v := range []int{v1, v2, v3, v4, v5} {
		fn(v)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: formatting_nbit

	   [Name] "formatting_nbit"
	   0000000001111111        0x007f    127
	   0000000100000000        0x0100    256
	   0000010000000000        0x0400   1024
	   0101010101010101        0x5555  21845
	   1010101010101010        0xaaaa  43690


	   [Elapsed] 39.22µs
	*/

}
