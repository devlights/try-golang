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
	//
	// 今回のような非同期処理を書いた場合、変数 data は競合状態が発生するので
	// 以下のような結果となる可能性がある
	//
	// - ゴルーチンが開始される前に if に到達してしまい、 data は 0
	// - ゴルーチンが終了した後に if に到達してしまい、 data は 1
	// - if が評価された時点では data は 0 だが、結果を出力する前にゴルーチンが動き data は 1
	if data == 0 {
		output.Stdoutl("[result]", "data is zero")
	} else {
		output.Stdoutl("[reslt]", "data is non-zero", data)
	}

	return nil
}
