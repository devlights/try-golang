// Stackoverflow Go Collective example
//
// How do you write multiline strings in Go?
//
// URL
//   - https://stackoverflow.com/questions/7933460/how-do-you-write-multiline-strings-in-go
//
// REFERENCES
//   - http://golang.org/doc/go_spec.html#String_literals
package main

import (
	"fmt"
	"strings"
)

func main() {
	// python では """ """ で multiline-strings が表現できるが
	// Go では ` ` で同じことになる

	var (
		multilineStrings = `this is 1st line.
this is 2nd line.
	this is 3rd line.
		EOF`
	)

	fmt.Println(multilineStrings)
	fmt.Printf("%v\n", strings.Split(multilineStrings, "\n"))
}
