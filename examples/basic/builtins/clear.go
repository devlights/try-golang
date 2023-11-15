package builtins

import "github.com/devlights/gomy/output"

// Clear は、Go 1.21 で追加されたビルトイン関数のclearについてのサンプルです.
func Clear() error {
	//
	// Go 1.21 から、clear ビルトイン関数が追加された
	// 対象となるのは、スライスとマップ。
	//   - スライスの場合は要素がゼロ値にクリアされる
	//   - マップの場合はキーが全て削除される
	//
	var (
		s = []int{1, 2, 3}
		m = map[int]string{100: "hello", 200: "world"}
		p = func(prefix string) {
			output.Stdoutf(prefix, "[slice] %v\t[map] %v\n", s, m)
		}
	)

	p("before")

	clear(s) // スライスの場合は要素がゼロ値になる。要素数は変わらない。
	clear(m) // マップの場合はキーがクリアされる。要素数はゼロとなる。

	p("after")

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: builtin_clear

	   [Name] "builtin_clear"
	   before               [slice] [1 2 3]    [map] map[100:hello 200:world]
	   after                [slice] [0 0 0]    [map] map[]


	   [Elapsed] 56.62µs
	*/

}
