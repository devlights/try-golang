package maths

import (
	"fmt"
	"math"
)

type mm struct {
	min, max interface{}
}

// MinMax -- MinMaxは各数値型の最小値と最大値を表示するサンプルです
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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: maths_minmax

	   [Name] "maths_minmax"
	   Int8 Min[-128] Max[127]
	   Int16 Min[-32768] Max[32767]
	   Int32 Min[-2147483648] Max[2147483647]
	   Int64 Min[-9223372036854775808] Max[9223372036854775807]


	   [Elapsed] 320.06µs
	*/

}
