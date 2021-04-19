package slices

import "github.com/devlights/gomy/output"

// NilAppend -- Nilなスライスに対して append した場合の挙動についてのサンプル
func NilAppend() error {
	// スライスのゼロ値はnil, len(nilスライス)は 0 となる.
	var ints []int

	output.Stdoutl("[ints == nil]", ints == nil)
	output.Stdoutl("[len(ints)]", len(ints))

	// Nil な スライスに対して append をしても問題なく追加できる
	ints = append(ints, []int{1, 2, 3, 4}...)

	output.Stdoutl("[ints]", ints)

	return nil
}
