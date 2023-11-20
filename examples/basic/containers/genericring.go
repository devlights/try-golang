package containers

import (
	"fmt"

	"github.com/devlights/try-golang/examples/basic/containers/generic/ring"

	"github.com/devlights/gomy/output"
)

// GenericRing は、container/ring/Ring をジェネリックにしたもののサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/container/ring@go1.21.4
func GenericRing() error {
	var (
		r = ring.New[string](5)
		n = r.Len()
	)

	// 値を設定（１）
	for i := 0; i < n; i++ {
		r.SetValue(fmt.Sprintf("helloworld-%d", i))
		r = r.Next()
	}

	// 現在のリングを出力
	r.Do(func(v string) {
		output.Stdoutl("[v]", v)
	})
	output.StdoutHr()

	// 再度値を設定（２）
	for i := 0; i < n; i++ {
		r.SetValue(fmt.Sprintf("helloworld-%d", i+100))
		r = r.Next()
	}

	// 現在のリングを出力
	r.Do(func(v string) {
		output.Stdoutl("[v]", v)
	})
	output.StdoutHr()

	// 飛び飛びで設定（３）
	r.SetValue("helloworld-997")
	r = r.Next()
	r = r.Next()
	r.SetValue("helloworld-998")
	r = r.Next()
	r.SetValue("helloworld-999")

	// 現在のリングを出力
	r.Do(func(v string) {
		output.Stdoutl("[v]", v)
	})

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: containers_generic_ring

	   [Name] "containers_generic_ring"
	   [v]                  helloworld-0
	   [v]                  helloworld-1
	   [v]                  helloworld-2
	   [v]                  helloworld-3
	   [v]                  helloworld-4
	   [v]                  helloworld-999
	   [v]                  helloworld-104
	   [v]                  helloworld-998
	   [v]                  helloworld-101
	   [v]                  helloworld-102


	   [Elapsed] 156.91µs
	*/

}
