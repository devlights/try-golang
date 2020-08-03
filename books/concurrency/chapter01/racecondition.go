package chapter01

import (
	"github.com/devlights/gomy/output"
)

// RaceCondition -- 競合状態のサンプルです.
//
// REFERENCES:
//   - P.5
func RaceCondition() error {
	var (
		data int
	)

	go func() {
		data++
	}()

	// ここで data の値は非決定的である
	// (0 かもしれないし 1 かもしれない)
	//
	// どの順番で処理が走るのかは保証されていない.
	if data == 0 {
		output.Stdoutl("[result]", "data is zero")
	}

	return nil
}
