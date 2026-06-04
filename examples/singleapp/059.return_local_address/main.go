package main

import "fmt"

func getvalue(x int) *int {
	return &x
}

func main() {
	p1 := getvalue(1)
	p2 := getvalue(2)
	fmt.Printf("%d,%d\n", *p1, *p2)
}
