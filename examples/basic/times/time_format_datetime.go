package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// FormatDateTime は、Go1.20で追加された time.DateTime フォーマット書式についてのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/time@go1.20.2#pkg-constants
func FormatDateTime() error {
	//
	// Go1.20 から、time.DateTime (yyyy-MM-dd HH:mm:ss) というフォーマットが追加された
	// これにより、少しだけフォーマットする際に楽になった
	//
	var (
		locJst *time.Location
		now    time.Time
		jst    time.Time
		err    error
	)

	locJst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	now = time.Now()
	jst = now.In(locJst)

	output.Stdoutf("[UTC          ]", "%v\n", now.UTC())
	output.Stdoutf("[JST          ]", "%v\n", jst)
	output.Stdoutf("[time.DateTime]", "%s\n", time.DateTime)
	output.Stdoutf("[time.Format  ]", "%s\n", jst.Format(time.DateTime))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_format_datetime

	   [Name] "time_format_datetime"
	   [UTC          ]      2024-03-27 06:03:37.823286141 +0000 UTC
	   [JST          ]      2024-03-27 15:03:37.823286141 +0900 JST
	   [time.DateTime]      2006-01-02 15:04:05
	   [time.Format  ]      2024-03-27 15:03:37


	   [Elapsed] 87.13µs
	*/

}
