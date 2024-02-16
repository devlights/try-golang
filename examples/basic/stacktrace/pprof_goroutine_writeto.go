package stacktrace

import (
	"bytes"
	"fmt"
	"io"
	"runtime/pprof"

	"github.com/devlights/gomy/output"
)

// PprofGoroutineWriteTo -- pprof.Lookup("goroutine").WriteTo() のサンプルです.
// REFERENCES
//   - https://pkg.go.dev/runtime/pprof#Lookup
//   - https://pkg.go.dev/runtime/pprof#Profile.WriteTo
//   - https://pkg.go.dev/runtime/pprof#Profile
//   - https://stackoverflow.com/questions/19094099/how-to-dump-goroutine-stacktraces
func PprofGoroutineWriteTo() error {
	const (
		profile           = "goroutine" // https://pkg.go.dev/runtime/pprof#Profile に事前定義プロファイルが記載されている
		useGoroutineStyle = 2           // goroutineの場合 2 を指定すると debug.Stack() と同じフォーマットになる.
	)

	var (
		errCh = make(chan error, 1)
		buf   = bytes.NewBuffer(nil)
		fn    = func(w io.Writer) error {
			p := pprof.Lookup(profile)
			if p == nil {
				return fmt.Errorf("profile does not exists (%s)", profile)
			}

			return p.WriteTo(w, useGoroutineStyle)
		}
	)

	go func() {
		defer close(errCh)
		if err := fn(buf); err != nil {
			errCh <- err
		}
	}()

	for e := range errCh {
		return e
	}

	output.Stdoutl("pprof", buf.String())

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: stacktrace_pprof_writeto

	   [Name] "stacktrace_pprof_writeto"
	   pprof                goroutine 6 [running]:
	   runtime/pprof.writeGoroutineStacks({0xa94620, 0xc0000bda70})
	           /home/gitpod/go/src/runtime/pprof/pprof.go:743 +0x6a
	   runtime/pprof.writeGoroutine({0xa94620?, 0xc0000bda70?}, 0x0?)
	           /home/gitpod/go/src/runtime/pprof/pprof.go:732 +0x25
	   runtime/pprof.(*Profile).WriteTo(0x9b5ab7?, {0xa94620?, 0xc0000bda70?}, 0x0?)
	           /home/gitpod/go/src/runtime/pprof/pprof.go:369 +0x14b
	   github.com/devlights/try-golang/examples/basic/stacktrace.PprofGoroutineWriteTo.func1({0xa94620, 0xc0000bda70})
	           /workspace/try-golang/examples/basic/stacktrace/pprof_goroutine_writeto.go:33 +0x46
	   github.com/devlights/try-golang/examples/basic/stacktrace.PprofGoroutineWriteTo.func2()
	           /workspace/try-golang/examples/basic/stacktrace/pprof_goroutine_writeto.go:39 +0x64
	   created by github.com/devlights/try-golang/examples/basic/stacktrace.PprofGoroutineWriteTo in goroutine 1
	           /workspace/try-golang/examples/basic/stacktrace/pprof_goroutine_writeto.go:37 +0x9c

	   goroutine 1 [chan receive]:
	   github.com/devlights/try-golang/examples/basic/stacktrace.PprofGoroutineWriteTo()
	           /workspace/try-golang/examples/basic/stacktrace/pprof_goroutine_writeto.go:44 +0xb1
	   github.com/devlights/try-golang/runner.(*Exec).Run(0x6070106?)
	           /workspace/try-golang/runner/exec.go:52 +0x131
	   github.com/devlights/try-golang/runner.(*Loop).exec(0xc00002e2d0?, {0x9c2a49, 0x18}, 0xc0000bda40)
	           /workspace/try-golang/runner/loop.go:126 +0x85
	   github.com/devlights/try-golang/runner.(*Loop).Run(0xc00007c410)
	           /workspace/try-golang/runner/loop.go:79 +0x23e
	   github.com/devlights/try-golang/cmd.Execute()
	           /workspace/try-golang/cmd/root.go:66 +0x612
	   main.main()
	           /workspace/try-golang/main.go:6 +0xf



	   [Elapsed] 149.21µs
	*/

}
