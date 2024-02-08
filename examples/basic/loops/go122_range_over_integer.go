package loops

import "github.com/devlights/gomy/output"

// Go122RangeOverInterger は、Go 1.22 で導入された range over integers ループ機能のサンプルです.
//
// 注意点として、通常 for-range ループでは インデックス と 値 を受け取るが
// range over integer を利用したループの場合は、インデックス のみとなる。（当然であるが）
//
// # REFERENCES
//
//   - https://go.dev/doc/go1.22#language
//   - https://go.dev/play/p/ky02zZxgk_r?v=gotip
//   - https://go.dev/ref/spec#For_range
func Go122RangeOverInterger() error {

	// for-range にて 数値 をそのまま指定できるようになった
	for i := range 3 {
		output.Stdoutl("[i]", i)
	}

	output.StdoutHr()

	var (
		count = 3
	)

	for i := range count {
		output.Stdoutl("[i]", i)
	}

	return nil

	/*
		$ task
		task: [build] go build .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: loops_go122_range_over_integer

		[Name] "loops_go122_range_over_integer"
		[i]                  0
		[i]                  1
		[i]                  2
		--------------------------------------------------
		[i]                  0
		[i]                  1
		[i]                  2


		[Elapsed] 43.88µs
	*/

}
