package loopiterator

import (
	"strconv"

	"github.com/devlights/gomy/output"
)

// CommonMistakePattern は、Goにてループ変数を扱う際によくある間違いについてのサンプルです.
//
// REFERENCES::
//   - https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
func CommonMistakePattern() error {
	// -------------------------------------------------------------------------
	// Goのループ変数は他の言語に無いクセがあり、ループ毎にループ変数を割り当てて
	// 処理するのではなく、ループ全体で一つの値を使い回すようになっている。
	// (変数のポインタは同じで値だけが変わっていく)
	//
	// なので、上記のWikiにあるようにポインタを保持するスライスなどを持っている状態で
	// ループ変数のポインタをそのまま格納するようなことをしてしまうと、最終的に全部同じ
	// データになってしまう。（ポインタが同じなので、ループの最終番目のデータになっている)
	//
	// 解決方法は、利用する前にループ変数の「コピー」をちゃんと取ること。
	// コピーを格納するようにしておけば、同じポインタになってしまうことは無い。
	// -------------------------------------------------------------------------
	bad()

	output.Stdoutl("--------------------------------------------------")

	good()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: using_ref_to_loop_iterator_variable

	   [Name] "using_ref_to_loop_iterator_variable"
	   0                    p=0xc0001a2908     v=0
	   1                    p=0xc0001a2910     v=1
	   2                    p=0xc0001a2918     v=2
	   3                    p=0xc0001a2920     v=3
	   4                    p=0xc0001a2928     v=4
	   --------------------------------------------------
	   0                    p=0xc0001a2938     v=0
	   1                    p=0xc0001a2940     v=1
	   2                    p=0xc0001a2948     v=2
	   3                    p=0xc0001a2950     v=3
	   4                    p=0xc0001a2958     v=4


	   [Elapsed] 57.61µs
	*/

}

func bad() {
	var (
		items []*int
	)

	for i := 0; i < 5; i++ {
		// コピーを取らずにループ変数のポインタを格納している
		// なので、結局全部同じものを格納していることになる
		items = append(items, &i)
	}

	for i, v := range items {
		output.Stdoutf(strconv.Itoa(i), "p=%p\tv=%v\n", v, *v)
	}
}

func good() {
	var (
		items []*int
	)

	for i := 0; i < 5; i++ {
		// 格納する前にループ変数のコピーを作っている
		// なので、ループ毎にちゃんと異なる値となる
		iCopy := i
		items = append(items, &iCopy)
	}

	for i, v := range items {
		output.Stdoutf(strconv.Itoa(i), "p=%p\tv=%v\n", v, *v)
	}
}
