package times

import (
	"fmt"
	"time"
)

// TimeParse は、 time.Parse() のサンプルです.
func TimeParse() error {
	// ------------------------------------------------------------
	// time.Parse(layout, value string) (Time, error)
	//
	// ref: https://golang.org/time/#Parse
	//
	// layout に指定した書式で、valueに指定された文字列を日付として解析する。
	// Goの日付書式は変わっていて、 yyyyとかの指定ではなく以下のように決まっている。
	//
	// 年： 2006
	// 月： 01
	// 日： 02
	// 時：　15
	// 分:  04
	// 秒:  05
	//
	// タイムゾーンは、 Z07:00 と指定する。
	//
	// つまり、日本でよく使う yyyy/MM/dd HH:mm:ss は以下のようにする。
	//
	//     2006/01/02 15:04:05
	//
	// タイムゾーンも付ける場合は、こうなる
	//
	//     2006/01/02 15:04:05Z07:00
	//
	// ------------------------------------------------------------
	// 正しい書式を与えると time.Time が取得できる
	d1, err := time.Parse(time.RFC3339, "2019-11-07T10:11:12+09:00")
	if err != nil {
		return err
	}

	fmt.Println("d1", d1)

	// 上で取得した d1 は、ローカルタイムなのでUTCにするには以下のようにする
	d2 := d1.UTC()
	fmt.Println("d2", d2)

	// 不正な値を指定するとエラーとなる
	d3, err := time.Parse(time.RFC3339, "not a date")
	if err != nil {
		fmt.Println(err)
	} else {
		// ここに来ることは無いが、d3の値は 「0001-01-01T00:00:00Z」 となる
		fmt.Println("d3", d3)
	}

	// 日本でよく使われる yyyy/MM/dd HH:mm:ss の場合はこのようになる
	const jpnFormat = "2006/01/02 15:04:05"
	d4, err := time.Parse(jpnFormat, "2019/11/07 10:11:12")
	if err != nil {
		return err
	}

	fmt.Println("d4", d4)

	// 上記でParseした d4 は、タイムゾーンを付与していないので UTC となる
	// JSTの状態でParseしたい場合は以下のようにする
	//
	// time.Local は、 ローカルタイムゾーンを表す
	// time.UTC は、 UTCタイムゾーンを表す
	d5, err := time.ParseInLocation(jpnFormat, "2019/11/07 10:11:12", time.Local)
	if err != nil {
		return err
	}

	fmt.Println("d5", d5)

	// または、フォーマットの方にタイムゾーンを付与して、Parseしても良い
	const jpnFormatWithTimeZone = "2006/01/02 15:04:05Z07:00"
	d6, err := time.Parse(jpnFormatWithTimeZone, "2019/11/07 10:11:12+09:00")
	if err != nil {
		return err
	}

	fmt.Println("d6", d6)

	// 無理やりだが、 UTC時間を9時間減算してタイムゾーン情報を当てても同じ結果となる
	fmt.Println(d4.Add(-9 * time.Hour).In(time.Local))

	// ローカルタイムゾーンに関しては、以下でも取得できる
	loc, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Println(d4.Add(-9 * time.Hour).In(loc))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_parse

	   [Name] "time_parse"
	   d1 2019-11-07 10:11:12 +0900 +0900
	   d2 2019-11-07 01:11:12 +0000 UTC
	   parsing time "not a date" as "2006-01-02T15:04:05Z07:00": cannot parse "not a date" as "2006"
	   d4 2019-11-07 10:11:12 +0000 UTC
	   d5 2019-11-07 10:11:12 +0000 UTC
	   d6 2019-11-07 10:11:12 +0900 +0900
	   2019-11-07 01:11:12 +0000 UTC
	   2019-11-07 10:11:12 +0900 JST


	   [Elapsed] 900.519µs
	*/

}
