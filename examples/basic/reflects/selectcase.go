package reflects

import (
	"reflect"

	"github.com/devlights/gomy/output"
)

// SelectCase -- reflect.SelectCase のサンプルです。
//
// REFERENCES:
//   - https://dev.to/hgsgtk/handling-with-arbitrary-channels-by-reflectselect-4d5g
func SelectCase() error {
	// 複数のチャネルを持っていて、その中のどれかのチャネルからでも良いので
	// データを取得していきたい場合、通常は for-select を使って処理することになる。
	// ただし、その場合、必要な分だけ case を増やしていかないと行けない.
	//
	// reflect.Select() を利用することで、リフレクションを利用して処理することが出来る

	const (
		chCount = 5
	)

	// チャネルを５つ用意
	chs := make([]chan int, chCount)
	for i := 0; i < chCount; i++ {
		ch := make(chan int)

		go func(i int) {
			ch <- i * i
		}(i)

		chs[i] = ch
	}

	// reflect.SelectCase を用意
	scs := make([]reflect.SelectCase, len(chs))
	for i, ch := range chs {
		sc := reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		}

		scs[i] = sc
	}

	// リフレクションを使って、複数のチャネルから一つ値を取得
	for i := 0; i < chCount; i++ {
		chosen, recv, ok := reflect.Select(scs)
		if ok {
			output.Stdoutf("reflect.Select", "chosen: %v\trecv: %v\n", chosen, recv)
		}
	}

	// 使ったチャネルを閉じる（このプログラムでは必要ないけど、お作法として）
	for _, ch := range chs {
		close(ch)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: reflect_selectcase

	   [Name] "reflect_selectcase"
	   reflect.Select       chosen: 4  recv: 16
	   reflect.Select       chosen: 2  recv: 4
	   reflect.Select       chosen: 3  recv: 9
	   reflect.Select       chosen: 0  recv: 0
	   reflect.Select       chosen: 1  recv: 1

	   [Elapsed] 232.21µs
	*/
}
