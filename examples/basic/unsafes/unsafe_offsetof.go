package unsafes

import (
	"fmt"
	"math"
	"unsafe"
)

// Offsetof は、unsafe.Offsetof() のサンプルです。
//
// > Offsetof returns the offset within the struct of the field represented by x, which must be of the form structValue.field.
// > In other words, it returns the number of bytes between the start of the struct and the start of the field.
// > The return value of Offsetof is a Go constant if the type of the argument x does not have variable size.
//
// > Offsetof は、引数 x によって表される構造体内のフィールドのオフセットを返します。引数 x は、必ず structValue.field の形式でなければなりません。
// > 言い換えると、これは構造体の先頭からそのフィールドの先頭までのバイト数を返します。
// > Offsetof の戻り値は、引数 x の型が可変サイズを持たない場合、Goの定数になります。
//
// REFERENCES:
//   - https://pkg.go.dev/unsafe@go1.25.3#Offsetof
func Offsetof() error {
	type (
		// わざとパディングが入る構造体とする
		// size: 24bytes
		S1 struct {
			A uint8  // offset=0,  size=1, padding=3
			B uint32 // offset=4,  size=4, padding=0
			C uint64 // offset=8,  size=8, padding=0
			D uint16 // offset=16, size=2, padding=6
		}
	)
	var (
		s1   = S1{math.MaxUint8, math.MaxUint32, math.MaxUint64, math.MaxUint16}
		off1 = unsafe.Offsetof(s1.A)
		off2 = unsafe.Offsetof(s1.B)
		off3 = unsafe.Offsetof(s1.C)
		off4 = unsafe.Offsetof(s1.D)
		size = unsafe.Sizeof(s1)
	)
	fmt.Printf("S1.A: %d\n", off1)
	fmt.Printf("S1.B: %d\n", off2)
	fmt.Printf("S1.C: %d\n", off3)
	fmt.Printf("S1.D: %d\n", off4)
	fmt.Printf("size: %d\n", size)

	return nil

	/*
		$ task
		task: [build] go build -o "/home/dev/dev/github/try-golang/try-golang" .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: unsafe_offsetof

		[Name] "unsafe_offsetof"
		S1.A: 0
		S1.B: 4
		S1.C: 8
		S1.D: 16
		size: 24


		[Elapsed] 11.773µs
	*/
}
