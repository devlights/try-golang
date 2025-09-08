package sliceop

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

// Reverse -- スライスのリバース処理についてのサンプルです。
func Reverse() error {

	var (
		ints = []int{
			1, 2, 3, 4, 5,
		}

		strs = []string{
			"hello", "world",
		}

		f = func(i []int, s []string) {
			fmt.Printf("[original]\tints[%v]\tstrs[%v]\n", i, s)
		}
	)

	f(ints, strs)

	reverseInts(ints)
	reverseStrs(strs)

	f(ints, strs)

	return nil

	/*
	    $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: slice_reverse

	   [Name] "slice_reverse"
	   [original]      ints[[1 2 3 4 5]]       strs[[hello world]]
	   [original]      ints[[5 4 3 2 1]]       strs[[world hello]]


	   [Elapsed] 18.43µs
	*/

}
