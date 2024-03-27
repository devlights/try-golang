package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// DaysInMonth は、月の日数を求めるサンプルです.
//
// # REFERENCES
//   - https://cs.opensource.google/go/go/+/refs/tags/go1.21.4:src/time/time.go;l=1467
func DaysInMonth() error {
	//
	// time.Date() で日の値を 0 で指定すると内部で正規化されて
	// -1日する動きとなる。なので、月を1加算して日を0にすると
	// その月の末日となる。
	//

	var (
		year   = time.Now().Year()
		months = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	)

	for _, m := range months {
		var (
			daysInMonth = time.Date(year, time.Month(m+1), 0, 0, 0, 0, 0, time.UTC).Day()
		)

		output.Stdoutf("[日数]", "%02d月の日数： %d\n", m, daysInMonth)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_daysinmonth

	   [Name] "time_daysinmonth"
	   [日数]                 01月の日数： 31
	   [日数]                 02月の日数： 29
	   [日数]                 03月の日数： 31
	   [日数]                 04月の日数： 30
	   [日数]                 05月の日数： 31
	   [日数]                 06月の日数： 30
	   [日数]                 07月の日数： 31
	   [日数]                 08月の日数： 31
	   [日数]                 09月の日数： 30
	   [日数]                 10月の日数： 31
	   [日数]                 11月の日数： 30
	   [日数]                 12月の日数： 31


	   [Elapsed] 137.97µs
	*/

}
