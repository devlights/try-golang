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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: panics_panic_and_defer

	   [Name] "panics_panic_and_defer"
	   [raise]              call defer
	   [caller]             call defer
	   [root]               call defer


	   [Elapsed] 50.4µs
	   END
	   panic: test

	   goroutine 1 [running]:
	   github.com/devlights/try-golang/examples/basic/panics.PanicAndDefer.func1()
	           /workspace/try-golang/examples/basic/panics/panic_and_defer.go:12 +0x68
	   github.com/devlights/try-golang/examples/basic/panics.PanicAndDefer.func2()
	           /workspace/try-golang/examples/basic/panics/panic_and_defer.go:16 +0x65
	   github.com/devlights/try-golang/examples/basic/panics.PanicAndDefer()
	           /workspace/try-golang/examples/basic/panics/panic_and_defer.go:20 +0x90
	   github.com/devlights/try-golang/runner.(*Exec).Run(0x6070100?)
	           /workspace/try-golang/runner/exec.go:52 +0x135
	   github.com/devlights/try-golang/runner.(*Loop).exec(0xc00002e2d0?, {0x9a7126, 0x16}, 0xc0000af950)
	           /workspace/try-golang/runner/loop.go:126 +0x85
	   github.com/devlights/try-golang/runner.(*Loop).Run(0xc000074408)
	           /workspace/try-golang/runner/loop.go:79 +0x23d
	   github.com/devlights/try-golang/cmd.Execute()
	           /workspace/try-golang/cmd/root.go:66 +0x627
	   main.main()
	           /workspace/try-golang/main.go:6 +0xf
	*/

}
