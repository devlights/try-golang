package panics

import "github.com/devlights/gomy/output"

// PanicAndDefer -- panicが呼ばれた場合でもdeferは処理されることを確認するサンプルです.
func PanicAndDefer() error {
	defer output.Stdoutl("[root]", "call defer")

	var (
		raise = func() {
			defer output.Stdoutl("[raise]", "call defer")
			panic("test")
		}
		proc = func() {
			defer output.Stdoutl("[caller]", "call defer")
			raise()
		}
	)

	proc()

	return nil
}
