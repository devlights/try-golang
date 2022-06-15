package unsafes

import (
	"fmt"
	"unsafe"
)

// Sizeof は、 unsafe.Sizeof() についてのサンプルです.
func Sizeof() error {
	// ---------------------------------------------------------
	// unsafe.Sizeof()
	//
	//   - Cのように 指定した 値の型メモリサイズを算出して返してくれる
	//   - 戻り値は uintptr 型
	//   - あくまでも指定した型から参照できる範囲のメモリサイズのみを返す
	//     - 例えばスライスの場合、 Sizeof が返すのはスライスのヘッダ部分のサイズ
	//     - スライスに実際格納されているデータのメモリサイズは含めてくれない
	// ---------------------------------------------------------
	type mem struct {
		b    bool
		i    int
		i16  int16
		i32  int32
		i64  int64
		s    string
		s2   string
		arr  [3]int
		sli  []int
		sli2 []int
	}

	m := mem{
		b:   true,
		i:   1,
		i16: 2,
		i32: 3,
		i64: 4,
		s:   "",
		s2:  "helloworld",
		arr: [3]int{1, 2, 3},
		sli: []int{},
		sli2: []int{
			1, 2, 3,
		},
	}

	fmt.Println("bool", unsafe.Sizeof(m.b))
	fmt.Println("int", unsafe.Sizeof(m.i))
	fmt.Println("int16", unsafe.Sizeof(m.i16))
	fmt.Println("int32", unsafe.Sizeof(m.i32))
	fmt.Println("int64", unsafe.Sizeof(m.i64))
	fmt.Println("string_empty", unsafe.Sizeof(m.s))
	fmt.Println("string_not_empty", unsafe.Sizeof(m.s2))
	fmt.Println("[3]int", unsafe.Sizeof(m.arr))
	fmt.Println("slice_empty", unsafe.Sizeof(m.sli))
	fmt.Println("slice_not_empty", unsafe.Sizeof(m.sli2))

	fmt.Println("----")
	fmt.Println("mem", unsafe.Sizeof(m))
	return nil
}
