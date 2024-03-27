package times

import (
	"time"

	"github.com/devlights/gomy/output"
	"github.com/devlights/gomy/times"
)

// TruncateHours -- time.Timeから時刻部分を除去するサンプルです.
//
// 時刻部分を除去する場合は、 time.Truncate() に 24時間 を指定する.
//
// REFERENCES:
//   - https://pkg.go.dev/time@latest#Time.Truncate
func TruncateHours() error {
	var (
		now       = time.Now()
		truncated = now.Truncate(24 * time.Hour)
		t1        = times.Formatter(now).YyyyMmddHHmmss()
		t2        = times.Formatter(truncated).YyyyMmddHHmmss()
	)

	output.Stdoutl("[now      ]", t1)
	output.Stdoutl("[truncated]", t2)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_truncate_hours

	   [Name] "time_truncate_hours"
	   [now      ]          2024/3/27 06:02:08
	   [truncated]          2024/3/27 00:00:00


	   [Elapsed] 54.96µs
	*/

}
