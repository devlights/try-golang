package sliceop

import "github.com/devlights/gomy/output"

// AppendSpecialBehavior は、append() を利用する際の特別な挙動に付いてのサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/builtin@go1.20.6#append
func AppendSpecialBehavior() error {
	//
	// append() を利用する際、以下のパターンは特別扱いで型が異なっているが動作する
	//   - byteスライスに対して文字列を追加
	//
	var (
		sl = make([]byte, 0)
	)

	sl = append(sl, []byte{0x61, 0x62, 0x63, 0x20}...) // OK
	sl = append(sl, "helloworld"...)                   // これもOK

	output.Stdoutf("[sl]", "%s/n", sl)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: slice_append_special_behavior

	   [Name] "slice_append_special_behavior"
	   [sl]                 abc helloworld/n

	   [Elapsed] 6.91µs
	*/

}
