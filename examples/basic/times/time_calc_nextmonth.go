package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// CalcNextMonth は、翌月の日付を求めるサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/time@go1.21.4
func CalcNextMonth() error {
	//
	// 日付が範囲外の存在しない値になったとき、Goはその値を存在する日付に変換する。
	// この変換を「正規化」という。なので、翌月を求める際に元の値によっては
	// 翌々月となってしまうことがある。
	//

	const (
		LOC = "Asia/Tokyo"
	)

	var (
		jst, _   = time.LoadLocation(LOC)
		original = time.Date(2023, 10, 31, 0, 0, 0, 0, jst)
	)

	output.Stdoutl("[original]", original.Format(time.RFC3339))

	//
	// 11月には31日が存在しないため
	// そのままAddDateすると、正規化されて12月02日となってしまう。
	//

	var (
		nextMonth1 = original.AddDate(0, 1, 1)
	)

	output.Stdoutl("[nextmonth1]", nextMonth1.Format(time.RFC3339))

	//
	// ちゃんと計算して求めるやり方 (書籍「実用 GO言語」)
	// 年月日以外の部分（時分秒とナノ秒）は面倒なので0固定
	//

	var (
		year, month, day = original.Date()                                // 元の日付
		first            = time.Date(year, month, 1, 0, 0, 0, 0, jst)     // 元の日付の月初
		year2, month2, _ = first.AddDate(0, 1, 0).Date()                  // 月初の値を使って翌月の年と月を取得
		nextMonth2       = time.Date(year2, month2, day, 0, 0, 0, 0, jst) // 元の日付の日を当てて翌月値を仮作成
	)

	if month2 != nextMonth2.Month() {
		// 正規化が発生して翌々月になってるので調整
		//   - ２月進めて１日戻すと翌月の末日になる
		nextMonth2 = first.AddDate(0, 2, -1)
	}

	output.Stdoutl("[nextmonth1]", nextMonth2.Format(time.RFC3339))

	return nil
}
