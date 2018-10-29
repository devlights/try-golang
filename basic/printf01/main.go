package main

import "fmt"

type MyType struct {
	i int
	v int
	x int
	y int
}

func main() {
	// fmt.Printf() は　C言語の printf 関数と
	// 同じような使い勝手を提供する。 書式については
	// 以下を参照。
	//     https://golang.org/pkg/fmt/#hdr-Printing
	messageFormat := "Hello %s\n"
	fmt.Printf(messageFormat, "World")

	messageFormat = "Hello %d\n"
	fmt.Printf(messageFormat, 100)

	messageFormat = "Hello %v\n"
	fmt.Printf(messageFormat, MyType{i: 100, v: 111, x: 222, y: 333})

	messageFormat = "Hello %+v\n"
	fmt.Printf(messageFormat, MyType{i: 100, v: 111, x: 222, y: 333})
}
