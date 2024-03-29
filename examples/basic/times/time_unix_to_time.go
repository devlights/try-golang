package times

import (
	"fmt"
	"time"
)

// TimeUnixToTime は、 time.Unix(sec, nsec) のサンプルです.
func TimeUnixToTime() error {
	// ------------------------------------------------------------
	// time.Unix(sec int64, nsec int64) Time
	//
	// ref: https://golang.org/time/#Unix
	//      https://mattn.kaoriya.net/software/lang/go/20130620173712.htm
	//
	// 指定されたUnix秒に基づく日付をローカルタイムで返す。
	// 引数 sec が秒数、 nsec は ナノ秒 を表す。
	//
	// 取得した Time データは、ローカルタイムとなっているので UTC 形式にするには
	// UTC() メソッドを呼び出す。
	// ------------------------------------------------------------
	var (
		sec  int64 = 100
		nsec int64 = 0
	)

	// unix秒を日付に変換
	localTime := time.Unix(sec, nsec)
	// UTC形式に変換
	utcTime := localTime.UTC()

	fmt.Printf("[local] %v\n[utc  ] %v\n", localTime, utcTime)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_unix_to_time

	   [Name] "time_unix_to_time"
	   [local] 1970-01-01 00:01:40 +0000 UTC
	   [utc  ] 1970-01-01 00:01:40 +0000 UTC


	   [Elapsed] 83.31µs
	*/

}
