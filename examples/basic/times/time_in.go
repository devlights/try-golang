package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// TimeIn -- time.In() の使い方のサンプルです。
//
// 日時自体は変更せずにタイムゾーンだけ変更する場合は, time.In() を利用します。
//
// REFERENCES
//   - https://zenn.dev/hsaki/articles/go-time-cheatsheet#unix%E6%99%82%E9%96%93%E3%81%8B%E3%82%89time.time%E5%9E%8B%E3%81%B8%E3%81%AE%E5%A4%89%E6%8F%9B---time.unix%E9%96%A2%E6%95%B0
func TimeIn() error {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	// time.In() は、時刻の値（Unix秒）は変えずにタイムゾーンだけ変更するメソッド
	now := time.Now()
	utc := now.In(time.UTC)
	tokyo := now.In(loc)

	output.Stdoutl("[now  ]", now, now.Unix())
	output.Stdoutl("[utc  ]", utc, utc.Unix())
	output.Stdoutl("[tokyo]", tokyo, tokyo.Unix())

	return nil
}