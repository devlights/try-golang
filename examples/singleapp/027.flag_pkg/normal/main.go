package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		a = flag.Int("a", 0, "")
		b = flag.String("b", "", "")
		c = flag.Bool("c", false, "")
	)

	flag.Parse()

	fmt.Printf("a=%v\tb=%v\tc=%v\n", *a, *b, *c)
}
