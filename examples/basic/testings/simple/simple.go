package simple

import (
	"errors"
	"math"
)

// ErrOverflow -- int32 の計算にてオーバーフローが発生した場合のエラーです。
var ErrOverflow = errors.New("int32 overflow")

// ErrUnderflow -- int32 の計算にてアンダーフローが発生した場合のエラーです。
var ErrUnderflow = errors.New("int32 underflow")

// Add は、指定された int32 を足し合わせた結果を返します。
func Add(ints ...int32) (int32, error) {
	if len(ints) == 0 {
		return 0, nil
	}

	total := int64(0)
	for _, i := range ints {
		t := total + int64(i)

		if t < math.MinInt32 {
			return int32(total), ErrUnderflow
		}

		if t > math.MaxInt32 {
			return int32(total), ErrOverflow
		}

		total += int64(i)
	}

	return int32(total), nil
}
