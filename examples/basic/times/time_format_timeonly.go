package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// FormatTimeOnly は、Go1.20で追加された time.TimeOnly フォーマット書式についてのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/time@go1.20.2#pkg-constants
func FormatTimeOnly() error {
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
	output.Stdoutf("[time.TimeOnly]", "%s\n", time.TimeOnly)
	output.Stdoutf("[time.Format  ]", "%s\n", jst.Format(time.TimeOnly))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_format_timeonly

	   [Name] "time_format_timeonly"
	   [UTC          ]      2024-03-27 06:04:10.143507636 +0000 UTC
	   [JST          ]      2024-03-27 15:04:10.143507636 +0900 JST
	   [time.TimeOnly]      15:04:05
	   [time.Format  ]      15:04:10


	   [Elapsed] 111.169µs
	*/

}
