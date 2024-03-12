package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// FormatMicrosecond は、time.Format() にてマイクロ秒を出力するサンプルです。
//
// マイクロ秒をフォーマットするには ".000000" とする。
// ドットを付けないとフォーマットされないので注意。
//
// # REFERENCES
//   - https://pkg.go.dev/time@go1.22.1#Time.Format
func FormatMicrosecond() error {
	var (
		microsec  = func() string { return time.Now().Format("05.000000") }
		wait100ms = func() { time.Sleep((100 * 1000) * time.Microsecond) }
	)

	output.Stdoutl("[1]", microsec())
	wait100ms()
	output.Stdoutl("[2]", microsec())
	wait100ms()
	output.Stdoutl("[3]", microsec())

	for range 5 {
		wait100ms()
	}

	output.Stdoutl("[4]", microsec())

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_format_microsecond

	   [Name] "time_format_microsecond"
	   [1]                  31.695992
	   [2]                  31.796184
	   [3]                  31.896460
	   [4]                  32.397545


	   [Elapsed] 701.582286ms
	*/

}
