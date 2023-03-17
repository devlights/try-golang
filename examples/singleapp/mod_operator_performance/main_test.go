// REFERENCE: https://www.reddit.com/r/golang/comments/11spdom/go_is_23_times_slower_than_js_in_a_similar_code/
package main

import "testing"

func BenchmarkModOperatorInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		get_primes_int(20_000)
	}
}

func BenchmarkModOperatorUInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		get_primes_uint32(20_000)
	}
}

func get_primes_int(limit int) []int {
	primes := []int{}
	for x := 2; x < limit; x++ {
		is_prime := true
		for y := 2; y < x; y++ {
			if x%y == 0 {
				is_prime = false
				break
			}
		}
		if is_prime {
			primes = append(primes, x)
		}
	}
	return primes
}

func get_primes_uint32(limit uint32) []uint32 {
	primes := []uint32{}
	for x := uint32(2); x < limit; x++ {
		is_prime := true
		for y := uint32(2); y < x; y++ {
			if x%y == 0 {
				is_prime = false
				break
			}
		}
		if is_prime {
			primes = append(primes, x)
		}
	}
	return primes
}
