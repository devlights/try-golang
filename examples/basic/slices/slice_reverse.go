package slices

import "fmt"

func reverseInts(s []int) {

	var (
		first = 0
		last  = len(s) - 1
	)

	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

func reverseStrs(s []string) {

	var (
		first = 0
		last  = len(s) - 1
	)

	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

// SliceReverse -- スライスのリバース処理についてのサンプルです。
func SliceReverse() error {

	var (
		ints = []int{
			1, 2, 3, 4, 5,
		}

		strs = []string{
			"hello", "world",
		}

		f = func(i []int, s []string) {
			fmt.Printf("[original]\tints[%v]\tstrs[%v]\n", ints, strs)
		}
	)

	f(ints, strs)

	reverseInts(ints)
	reverseStrs(strs)

	f(ints, strs)

	return nil
}
