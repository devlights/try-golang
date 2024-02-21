package stdout

import "fmt"

// Printf01 -- 標準出力についてのサンプル
func Printf01() error {
	message := "Hello Golang!!"
	fmt.Printf("%s", message)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: printf01

	   [Name] "printf01"
	   Hello Golang!!

	   [Elapsed] 2.88µs
	*/

}
