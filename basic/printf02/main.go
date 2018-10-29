package main

import "fmt"

type MyData struct {
	age  int
	name string
}

func (self *MyData) GetProfile() string {
	return fmt.Sprintf("%v", self)
}

func (self *MyData) GetProfile2() string {
	return fmt.Sprintf("%#v", self)
}

func (self *MyData) GetProfile3() string {
	return fmt.Sprintf("%T", self)
}

func main() {
	data := MyData{age: 30, name: "hello world"}

	fmt.Println(data.GetProfile())
	fmt.Println(data.GetProfile2())
	fmt.Println(data.GetProfile3())
}
