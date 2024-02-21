package stdout

import "fmt"

// Println01 -- 標準出力についてのサンプル
func Println01() error {
	hello := "hello"
	world := "world"

	// fmt.Println() は、複数の値を指定するとスペースで区切って表示してくれる
	fmt.Println(hello, world)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: println01

	   [Name] "println01"
	   hello world


	   [Elapsed] 4.71µs
	*/

}
