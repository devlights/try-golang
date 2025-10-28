package unsafes

import (
	"bytes"
	"fmt"
	"unsafe"
)

// Slice は、unsafe.SliceData() と unsafe.Slice() のサンプルです。
//
// unsafe.SliceData() は、特定のスライスを *T に変換する関数。
// 逆を行ってくれるのが unsafe.Slice() となる。
//
// REFERENCES:
//   - https://pkg.go.dev/unsafe@go1.25.3#SliceData
func Slice() error {
	var (
		original = []byte("helloworld")
		result   []byte
		ptr      *byte
	)
	ptr = unsafe.SliceData(original)          // []byteを*byteに変換
	result = unsafe.Slice(ptr, len(original)) // *byteを[]byteに変換

	fmt.Printf("b1 equals b2: %v\n", bytes.Equal(original, result))

	return nil
}
