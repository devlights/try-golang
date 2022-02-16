package array

import "github.com/devlights/gomy/output"

// Ellipses -- 配列を ... で初期化するサンプルです.
func Ellipses() error {
	var (
		items1 = [5]int{1, 2, 3, 4, 5}
		items2 = [...]int{1, 2, 3, 4, 5} // 初期化時に ... を指定することもできる
		ch     = make(chan int)
	)

	go func(ch chan<- int) {
		defer close(ch)
		for _, v := range items1 {
			ch <- v
		}
		for _, v := range items2 {
			ch <- v
		}
	}(ch)

	for v := range ch {
		output.Stdoutl("v", v)
	}

	return nil
}
