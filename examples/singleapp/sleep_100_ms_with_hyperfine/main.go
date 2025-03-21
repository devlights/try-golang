package main

import (
	"flag"
	"time"
)

func main() {
	var d int

	flag.IntVar(&d, "d", 100, "")
	flag.Parse()

	time.Sleep(time.Duration(d) * time.Millisecond)
}
