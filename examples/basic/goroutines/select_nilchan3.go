package goroutines

import "github.com/devlights/gomy/output"

// SelectNilChan3 -- select ステートメントで nil チャネル を使って選択されるチャネルの有効・無効を切り替えるサンプルです (3).
//
// シンプルな形のサンプル。
func SelectNilChan3() error {
	var (
		gen = func(out chan<- int) {
			defer close(out)
			for i := 0; i < 5; i++ {
				out <- (i + 1)
			}
		}
		output = func(done chan<- any, in1, in2 <-chan int) {
			defer close(done)

		LOOP:
			for {
				select {
				case v, ok := <-in1:
					if !ok {
						in1 = nil
						output.Stderrl("[in1]", "close")
						continue
					}
					output.Stderrl("[in1]", v)
				case v, ok := <-in2:
					if !ok {
						in2 = nil
						output.Stderrl("[in2]", "close")
						continue
					}
					output.Stderrl("[in2]", v)
				default:
					if in1 == nil && in2 == nil {
						break LOOP
					}
				}
			}
		}
	)

	var (
		ch1  = make(chan int)
		ch2  = make(chan int)
		done = make(chan any)
	)

	go gen(ch1)
	go gen(ch2)
	go output(done, ch1, ch2)

	<-done

	return nil
}
