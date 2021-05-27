package loopiterator

import (
	"sync"

	"github.com/devlights/gomy/output"
)

type (
	value struct {
		i int
	}
)

// PassingLoopVariableToGoroutineByPointer -- ループ変数をポインタ経由でGoroutineに渡した場合のサンプルです.
//
// REFERENCES:
//   - https://stackoverflow.com/a/23637430
//   - https://golang.org/doc/effective_go.html#channels
func PassingLoopVariableToGoroutineByPointer() error {
	// ----------------------------------------------------------------
	// Go では、ループ変数は使いまわしされるので
	// 例えば、ループ内でGoroutineを起動する際にループ変数を
	// ポインタで渡すような書き方をすると、実際には同じアドレスを
	// 渡していることになる。
	//
	// なので、各Goroutineは、ほぼ同じ値を見ることになってしまう。
	// (Goroutineがとても早く起動して、変化する前の値をみた場合はその時の値が見える)
	//
	// 回避策としては
	// 渡す前にコピーをとって渡すようにするか、添字を使って
	// 値を取得するようにすれば大丈夫
	// ----------------------------------------------------------------
	badpattern()

	output.StdoutHr()

	goodpattern()

	return nil
}

func badpattern() {
	var (
		wg   = &sync.WaitGroup{}
		vals = []value{
			{1},
			{2},
			{3},
		}
	)

	wg.Add(len(vals))
	for _, v := range vals {
		output.Stdoutf("[bad][v]", "addr=%p, value=%v\n", &v, v)

		go func(v *value) {
			defer wg.Done()
			output.Stdoutf("[bad][v][goroutine]", "addr=%p, value=%v\n", v, *v)
		}(&v)
	}

	wg.Wait()
}

func goodpattern() {
	var (
		wg   = &sync.WaitGroup{}
		vals = []value{
			{1},
			{2},
			{3},
		}
	)

	wg.Add(len(vals))
	for _, v := range vals {
		output.Stdoutf("[good][v]", "addr=%p, value=%v\n", &v, v)

		// 渡す前にコピーを取得
		//   - https://stackoverflow.com/a/23637430
		v := v
		go func(v *value) {
			defer wg.Done()
			output.Stdoutf("[good][v][goroutine]", "addr=%p, value=%v\n", v, *v)
		}(&v)
	}

	wg.Wait()
}
