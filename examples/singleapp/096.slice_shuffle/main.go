package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// rand.Shuffle()を使うとスライスなどをシャッフル出来る

	var (
		s  = []byte("hello world")
		fn = func(i, j int) {
			s[i], s[j] = s[j], s[i]
		}
	)
	fmt.Println(string(s))
	rand.Shuffle(len(s), fn)
	fmt.Println(string(s))
}
