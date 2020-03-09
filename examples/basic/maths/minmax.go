package maths

import (
	"fmt"
	"math"
)

type mm struct {
	min, max interface{}
}

// MinMaxは各数値型の最小値と最大値を表示するサンプルです
func MinMax() error {

	// 各数値型の最小値と最大値は math パッケージに定義されている
	// Floatのみ、MinXXではなくて SmallestNonzeroFloatXX となる
	m := map[string]*mm{
		"Int8":  {min: math.MinInt8, max: math.MaxInt8},
		"Int16": {min: math.MinInt16, max: math.MaxInt16},
		"Int32": {min: math.MinInt32, max: math.MaxInt32},
		"Int64": {min: math.MinInt64, max: math.MaxInt64},
	}

	for k, v := range m {
		fmt.Printf("%v Min[%v] Max[%v]\n", k, v.min, v.max)
	}

	return nil
}
