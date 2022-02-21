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
}
