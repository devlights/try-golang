package structs

import (
	"fmt"
	"strconv"

	"github.com/devlights/gomy/output"
)

type (
	helloworld string
	worldhello byte
	myint      int
)

func (h helloworld) say() string {
	return fmt.Sprintf("hello world %s", h)
}

func (w worldhello) say() string {
	return fmt.Sprintf("world hello %d", w)
}

func (m myint) say() string {
	return strconv.Itoa(int(m))
}

// SameMethodOnEachTypes -- レシーバの型が異なる同名メソッド定義のサンプルです
func SameMethodOnEachTypes() error {
	var (
		h = helloworld("-- golang")
		w = worldhello(255)
		i = myint(100)
	)

	for _, v := range []interface{ say() string }{h, w, i} {
		p(v)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: struct_same_method

	   [Name] "struct_same_method"
	   [say]                [structs.helloworld  ]     hello world -- golang
	   [say]                [structs.worldhello  ]     world hello 255
	   [say]                [structs.myint       ]     100


	   [Elapsed] 39.289µs
	*/

}

func p(g interface{ say() string }) {
	output.Stdoutf("[say]", "[%-20T]\t%s\n", g, g.say())
}
