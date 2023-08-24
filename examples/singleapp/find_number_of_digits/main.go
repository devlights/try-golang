package main

import (
	"fmt"

	"github.com/devlights/try-golang/examples/singleapp/find_number_of_digits/lib"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	v := 123456789
	fmt.Printf("[ToString] %v\t[Log10] %v\n", lib.UseToString(v), lib.UseLog10(v))

	return nil
}