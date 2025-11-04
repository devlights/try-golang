package unsafes

import (
	"fmt"
	"math"
	"unsafe"
)

// Alignof は、unsafe.Alignof() のサンプルです。
//
// > Alignof takes an expression x of any type and returns the required alignment of a hypothetical variable v as if v was declared via var v = x.
// > It is the largest value m such that the address of v is always zero mod m.
// > It is the same as the value returned by reflect.TypeOf(x).Align().
// > As a special case, if a variable s is of struct type and f is a field within that struct, then Alignof(s.f) will return the required alignment of a field of that type within a struct.
// > This case is the same as the value returned by reflect.TypeOf(s.f).FieldAlign().
// > The return value of Alignof is a Go constant if the type of the argument does not have variable size.
//
// > Alignof は任意の型の式 x を受け取り、var v = x のように宣言されたと仮定した場合の、仮想的な変数 v に必要な**アライメント（配置要求）**を返します。
// > この戻り値は、v のアドレスが常に m で割って余りが 0 となるような最大の正の整数 m です。これは、reflect.TypeOf(x).Align() が返す値と同じです。
// > 特別なケースとして、変数 s が構造体型で、f がその構造体内のフィールドである場合、Alignof(s.f) は構造体内のその型のフィールドに必要なアライメントを返します
// > このケースは reflect.TypeOf(s.f).FieldAlign() が返す値と同じです。
// > 引数の型が可変サイズを持たない場合、Alignof の戻り値はGoの定数になります。
//
// REFERENCES:
//   - https://pkg.go.dev/unsafe@go1.25.3#Alignof
func Alignof() error {
	type (
		S1 struct {
			A uint8  // offset=0,  size=1, padding=3
			B uint32 // offset=4,  size=4, padding=0
			C uint64 // offset=8,  size=8, padding=0
			D uint16 // offset=16, size=2, padding=6
		}
		S2 struct {
			A uint8  // offset=0, size=1, padding=1
			B uint16 // offset=2, size=2, padding=0
			C uint8  // offset=4, size=1, padding=0
			D uint8  // offset=5, size=1, padding=0
		}
	)
	var (
		s1 = S1{math.MaxUint8, math.MaxUint32, math.MaxUint64, math.MaxUint16}
		s2 = S2{math.MaxUint8, math.MaxUint16, math.MaxUint8, math.MaxUint8}
	)
	fmt.Printf("S1 Alignment Size: %d\n", unsafe.Alignof(s1)) // C言語と同様に構造体内での最大アライメント要件を持つ型のサイズとなる (uint64)
	fmt.Printf("S1 Size          : %d\n", unsafe.Sizeof(s1))
	fmt.Printf("S2 Alignment Size: %d\n", unsafe.Alignof(s2)) // C言語と同様に構造体内での最大アライメント要件を持つ型のサイズとなる (uint16)
	fmt.Printf("S2 Size          : %d\n", unsafe.Sizeof(s2))

	return nil
}
