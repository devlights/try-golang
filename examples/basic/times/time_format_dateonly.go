package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// FormatDateOnly は、Go1.20で追加された time.DateOnly フォーマット書式についてのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/time@go1.20.2#pkg-constants
func FormatDateOnly() error {
	//
	// Go1.20 から、time.DateOnly (yyyy-MM-dd) というフォーマットが追加された
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
	output.Stdoutf("[time.DateOnly]", "%s\n", time.DateOnly)
	output.Stdoutf("[time.Format  ]", "%s\n", jst.Format(time.DateOnly))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_format_dateonly

	   [Name] "time_format_dateonly"
	   [UTC          ]      2024-03-27 06:03:54.384903008 +0000 UTC
	   [JST          ]      2024-03-27 15:03:54.384903008 +0900 JST
	   [time.DateOnly]      2006-01-02
	   [time.Format  ]      2024-03-27


	   [Elapsed] 153.36µs
	*/

}
