package main

func Plus[T ~int](x, y T) T {
	if x == 100 {
		return y
	}

	return x + y
}
