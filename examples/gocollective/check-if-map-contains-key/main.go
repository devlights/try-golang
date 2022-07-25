// Stackoverflow Go Collective example
//
// How to check if a map contains a key in Go?
//
// URL
//   - https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
//
// REFERENCES
package main

import "fmt"

func main() {
	// Goでmapに対象のキーがあるのかを確認するのは以下のようにする

	var (
		m = make(map[string]string)
	)
	m["hello"] = "world"

	v, ok := m["hello"]
	if !ok {
		fmt.Printf("key '%s' does not exists.", "hello")
	}
	fmt.Println(v)

	v, ok = m["world"]
	if !ok {
		fmt.Printf("key '%s' does not exists.", "world")
	}
	fmt.Println(v)
}
