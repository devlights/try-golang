package times

import (
	"fmt"
	"time"
)

// TimeNow は、 time.Now() のサンプルです.
func TimeNow() error {
	// ------------------------------------------------------------
	// time.Now() Time
	//
	// ref: https://golang.org/time/#Now
	//
	// 現在日時（ローカルタイム）を返す。
	// unix秒を取得したい場合は、さらに Unix() メソッドを呼び出す。
	// UTC時間を取得したい場合は、さらに UTC() メソッドを呼び出す。
	// フォーマットして表示する場合は Format() メソッドを呼び出す.
	// ------------------------------------------------------------
	nowLocalTime := time.Now()
	nowUtcTime := nowLocalTime.UTC()
	nowUnixSec := nowLocalTime.Unix()

	localTime := nowLocalTime.Format(time.RFC3339)
	utcTime := nowUtcTime.Format(time.RFC3339)

	fmt.Printf("local time: %v\nutc time:%v\nunix sec:%v\n", localTime, utcTime, nowUnixSec)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_now

	   [Name] "time_now"
	   local time: 2024-03-27T05:59:08Z
	   utc time:2024-03-27T05:59:08Z
	   unix sec:1711519148


	   [Elapsed] 74.83µs
	*/

}
