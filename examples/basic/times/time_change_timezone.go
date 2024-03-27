package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// ChangeTimeZone は、time.Timeをいろいろなタイム・ゾーンの値に変換するサンプルです.
func ChangeTimeZone() error {
	// JSTの現在時刻を取得
	locJst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	jst := time.Now().In(locJst)
	output.Stdoutf("[JST]", "%v\n", jst)

	// UTCへ変換
	utc := jst.UTC()
	output.Stdoutf("[UTC]", "%v\n", utc)

	// UTCからPDTに変換
	// (*) 夏時間(太平洋夏時間, Daylight Saving Time, DST)が適用されている場合でも
	// このコードで自動的にDSTが考慮される
	locPdt, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return err
	}

	pdt := utc.In(locPdt)
	output.Stdoutf("[PDT]", "%v\n", pdt)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_change_timezone

	   [Name] "time_change_timezone"
	   [JST]                2024-03-27 15:03:17.210910889 +0900 JST
	   [UTC]                2024-03-27 06:03:17.210910889 +0000 UTC
	   [PDT]                2024-03-26 23:03:17.210910889 -0700 PDT


	   [Elapsed] 659.019µs
	*/

}
