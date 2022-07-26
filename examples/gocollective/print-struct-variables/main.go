// Stackoverflow Go Collective example
//
// How to print struct variables in console
//
// URL
//   - https://stackoverflow.com/questions/24512112/how-to-print-struct-variables-in-console
//
// REFERENCES
//   - https://pkg.go.dev/fmt@latest
package main

import "fmt"

func main() {
	o := struct {
		Id    int
		Value string
	}{1, "hello world"}

	fmt.Printf("[%%v ] %v\n", o)
	fmt.Printf("[%%+v] %+v\n", o)

	fn(&o)
}

func fn(o *struct {
	Id    int
	Value string
}) {
	fmt.Printf("%+v\n", *o)
}
