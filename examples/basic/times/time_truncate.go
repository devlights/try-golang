package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// Truncate -- time.Truncate() のサンプルです。n分置き や n時間置き の時間を表現することができます。
//
// REFERENCES
//   - https://zenn.dev/mltokky/articles/20220426_golang_time_truncate
func Truncate() error {
	const (
		layout = "15:04"
	)

	var (
		values = []string{
			"10:00",
			"10:01",
			"10:15",
			"10:29",
			"10:30",
			"10:59",
			"11:00",
		}
	)

	output.Stdoutf("[truncate]", "%s\t%s\n", "Original", "Truncated")
	for _, v := range values {
		t, err := time.Parse(layout, v)
		if err != nil {
			return err
		}

		// 30分単位の時刻に切り捨て
		truncated := t.Truncate(30 * time.Minute)

		output.Stdoutf("[truncate]", "%s\t%s\n", v, truncated.Format(layout))
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_truncate

	   [Name] "time_truncate"
	   [truncate]           Original   Truncated
	   [truncate]           10:00      10:00
	   [truncate]           10:01      10:00
	   [truncate]           10:15      10:00
	   [truncate]           10:29      10:00
	   [truncate]           10:30      10:30
	   [truncate]           10:59      10:30
	   [truncate]           11:00      11:00


	   [Elapsed] 44.13µs
	*/

}
