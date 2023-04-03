package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// FormatDateTime は、Go1.20で追加された time.TimeOnly フォーマット書式についてのサンプルです.
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
}
