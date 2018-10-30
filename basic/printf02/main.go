package main

import "fmt"

func main() {
	data := MyData{30, "hello world"}

	fmt.Println(data.GetValue())
	fmt.Println(data.GetValue2())
	fmt.Println(data.GetValue3())
	fmt.Println(data.GetValue4())
}
