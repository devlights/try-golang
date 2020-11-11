package tutorial05

import "fmt"

// BasicTypes は、 Tour of Go - Basic types (https://tour.golang.org/basics/11) の サンプルです。
func BasicTypes() error {
	// ------------------------------------------------------------
	// Go言語の基本型は以下
	//
	// - bool
	// - string
	// - int     (int8,   int16,  int32,  int64)
	// - uint    (uint8 , uint16, uint32, uint64, uintptr)
	// - byte    (uint8 の 別名)
	// - rune    (int32 の 別名)
	// - float   (float32, float64)
	// - complex (complex32, complex64)
	// ------------------------------------------------------------
	//noinspection GoVarAndConstTypeMayBeOmitted
	var (
		boolVal bool   = true
		maxInt  int64  = 1<<63 - 1
		maxUInt uint64 = 1<<64 - 1
		byteVal byte   = 42
		runeVal rune   = 'a'
	)

	//noinspection GoBoolExpressions
	fmt.Println(boolVal, maxInt, maxUInt, byteVal, runeVal)

	return nil
}
