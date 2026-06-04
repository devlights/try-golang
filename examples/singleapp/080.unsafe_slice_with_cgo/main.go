package main

/*
extern void c_func();
*/
import "C"

func main() {
	C.c_func()
}

/*
	$ task
	task: [default] go run *.go
	[Go] helloworld
	[C ] dlrowolleh
*/
