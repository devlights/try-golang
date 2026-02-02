package main

import (
	"flag"
	"fmt"
)

func main() {
	var x int
	flag.IntVar(&x, "x", 0, "")
	flag.Parse()

	for i := range x {
		fmt.Println(i)
	}
}
