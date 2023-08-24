package lib

import (
	"math"
	"strconv"
)

type (
	NumberOfDigits int
)

func UseToString(v int) NumberOfDigits {
	return NumberOfDigits(len(strconv.Itoa(v)))
}

func UseLog10(v int) NumberOfDigits {
	return NumberOfDigits(int(math.Floor(math.Log10(float64(v)))) + 1)
}
