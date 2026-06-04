// REFERENCE: https://www.reddit.com/r/golang/comments/11spdom/go_is_23_times_slower_than_js_in_a_similar_code/
package main

import "testing"

func BenchmarkModOperatorInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getPrimesInt(20_000)
	}
}

func BenchmarkModOperatorUInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getPrimesUInt32(20_000)
	}
}

func getPrimesInt(limit int) []int {
	var primes []int
	for x := 2; x < limit; x++ {
		isPrime := true
		for y := 2; y < x; y++ {
			if x%y == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, x)
		}
	}
	return primes
}

func getPrimesUInt32(limit uint32) []uint32 {
	var primes []uint32
	for x := uint32(2); x < limit; x++ {
		isPrime := true
		for y := uint32(2); y < x; y++ {
			if x%y == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, x)
		}
	}
	return primes
}
