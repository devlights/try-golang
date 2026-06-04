package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		flags = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		a     = flags.Int("a", 0, "")
		b     = flags.String("b", "", "")
		c     = flags.Bool("c", false, "")
	)

	flags.Parse(os.Args[1:])

	fmt.Printf("a=%v\tb=%v\tc=%v\n", *a, *b, *c)
}
