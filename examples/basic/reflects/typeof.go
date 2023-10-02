package reflects

import (
	"reflect"

	"github.com/devlights/gomy/output"
)

// TypeOf -- reflect.TypeOf() のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/reflect@go1.21.1
//   - https://go.dev/blog/laws-of-reflection
func TypeOf() error {
	//
	// Goでリフレクションを利用する場合は reflect パッケージを使う。
	// 使い方としてはC#などの他の言語と同じイメージ。
	//
	// - 型を取得するには reflect.TypeOf()  を使用する
	// - 値を取得するには reflect.ValueOf() を使用する
	//

	type (
		MyInt int
		MySt  struct {
			s   string
			i   int64
			ui  uint64
			f   float64
			ch  chan string
			sli []int
			m   map[int]int
		}
	)

	var (
		i  = 0
		mi = MyInt(0)
		st = MySt{"helloworld", int64(100), uint64(100), float64(100.0), make(chan string), []int{1}, map[int]int{1: 2}}
	)

	var (
		ti  = reflect.TypeOf(i)
		tmi = reflect.TypeOf(mi)
		tst = reflect.TypeOf(st)
	)

	output.Stdoutf("[ti ]", "%T\t%v\t%dbytes\n", ti, ti, ti.Size())
	output.Stdoutf("[tmi]", "%T\t%v\t%dbytes\n", tmi, tmi, tmi.Size())
	output.Stdoutf("[tst]", "%T\t%v\t%dbytes\n", tst, tst, tst.Size())

	for idx := 0; idx < tst.NumField(); idx++ {
		f := tst.Field(idx)
		output.Stdoutf(" >>>", "\t%T:\t%s\t%dbytes\n", f, f.Type.Name(), f.Type.Size())

		// スライス、マップ、チャネルは f.Type.Name() では表示されない
	}

	return nil
}
