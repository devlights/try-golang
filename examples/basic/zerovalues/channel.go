package zerovalues

import "github.com/devlights/gomy/output"

// Channel は、Goにおける チャネル のゼロ値についてのサンプルです.
//
// # REFERENCES
//   - https://go.dev/tour/basics/12
//   - https://brain2life.hashnode.dev/default-zero-values-in-go
func Channel() error {
	//
	// チャネル の ゼロ値 は nil
	//
	var (
		ch chan int
	)

	output.Stdoutf("[chan zerovalue]", "%v\n", ch)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: zerovalues_chan

	   [Name] "zerovalues_chan"
	   [chan zerovalue]     <nil>


	   [Elapsed] 20.84µs
	*/

}
