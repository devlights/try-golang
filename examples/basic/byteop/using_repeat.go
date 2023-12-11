package byteop

import (
	"bytes"

	"github.com/devlights/gomy/output"
)

// UsingRepeat は、bytes.Repeat() のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/bytes@go1.21.5#Repeat
func UsingRepeat() error {
	//
	// インデックス付きの for ループを書くのが面倒なときに結構便利
	//
	for i, v := range bytes.Repeat([]byte{1}, 10) {
		output.Stdoutl("[i, v]", i, v)
	}

	output.StdoutHr()

	//
	// 所定のスライスを繰り返して処理したいとき
	//
	var (
		ch   = make(chan int)
		done = make(chan struct{})
	)

	go func(ch chan<- int) {
		defer close(ch)

		for _, v := range bytes.Repeat([]byte{1, 2, 3, 4, 5}, 3) {
			i := int(v)
			ch <- i * i
		}
	}(ch)

	go func(done chan<- struct{}, ch <-chan int) {
		defer close(done)

		for v := range ch {
			output.Stdoutl("[v]", v)
		}
	}(done, ch)

	<-done

	return nil
}
