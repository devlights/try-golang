package time_

import (
	"fmt"
	"time"
)

// TimeParse は、 time.Parse() のサンプルです.
func TimeParse() error {
	// ------------------------------------------------------------
	// time.Parse(layout, value string) (Time, error)
	//
	// ref: https://golang.org/pkg/time/#Parse
	//
	// layout に指定した書式で、valueに指定された文字列を日付として解析する。
	// ------------------------------------------------------------
	// 正しい書式を与えると time.Time が取得できる
	d1, err := time.Parse(time.RFC3339, "2019-11-07T10:11:12+09:00")
	if err != nil {
		return err
	}

	d2 := d1.UTC()

	fmt.Printf("d1: %s\nd2: %s\n", d1, d2)

	// 不正な値を指定するとエラーとなる
	d3, err := time.Parse(time.RFC3339, "not a date")
	if err != nil {
		fmt.Println(err)
	} else {
		// ここに来ることは無いが、d3の値は 「0001-01-01T00:00:00Z」 となる
		fmt.Println(d3)
	}

	return nil
}
