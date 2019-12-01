package struct_

import (
	"fmt"
	"unsafe"
)

// EmptyStruct は、空の構造体についサンプルです.
func EmptyStruct() error {
	// --------------------------------------------------
	// 空の構造体 (Empty struct)
	//   - 空のインターフェースはメモリサイズが 0
	//   - 型のメモリサイズは Unsage.sizeof() で調べられる
	//
	// ref: https://dave.cheney.net/2014/03/25/the-empty-struct
	// --------------------------------------------------
	var (
		emptyStruct    struct{}
		emptyInterface interface{}
	)

	emptyStruct = struct{}{}
	emptyInterface = emptyStruct

	s1 := unsafe.Sizeof(emptyStruct)
	s2 := unsafe.Sizeof(emptyInterface)
	fmt.Printf("EmptyStruct[%d] EmptyInterface[%d]\n", s1, s2)

	return nil
}
