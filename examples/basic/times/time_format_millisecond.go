package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// FormatMillisecond は、time.Format() にてミリ秒を出力するサンプルです。
//
// ミリ秒をフォーマットするには ".000" とする。
// ドットを付けないとフォーマットされないので注意。
//
// # REFERENCES
//   - https://pkg.go.dev/time@go1.22.1#Time.Format
func FormatMillisecond() error {
	var (
		millisec  = func() string { return time.Now().Format("05.000") }
		wait100ms = func() { time.Sleep(100 * time.Millisecond) }
	)

	output.Stdoutl("[1]", millisec())
	wait100ms()
	output.Stdoutl("[2]", millisec())
	wait100ms()
	output.Stdoutl("[3]", millisec())

	for range 5 {
		wait100ms()
	}

	output.Stdoutl("[4]", millisec())

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_format_millisecond

	   [Name] "time_format_millisecond"
	   [1]                  25.065
	   [2]                  25.165
	   [3]                  25.265
	   [4]                  25.766


	   [Elapsed] 701.858467ms
	*/

}
