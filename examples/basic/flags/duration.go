package flags

import (
	"flag"
	"time"

	"github.com/devlights/gomy/output"
)

// Duration は、flag.Duration(), flag.DurationVar() のサンプルです。
//
// flagパッケージの関数は、flag.Duration()のように受け皿を戻り値で返してくれる関数と
// flag.DurationVar() のように予め自前で用意している変数を利用する２パターンの使い方がある。
//
// # REFERENCES
//   - https://pkg.go.dev/flag@go1.22.4#Duration
//   - https://pkg.go.dev/flag@go1.22.4#DurationVar
func Duration() error {
	var (
		fs = flag.NewFlagSet("", flag.ExitOnError)

		d1 *time.Duration
		d2 time.Duration
	)

	d1 = fs.Duration("d1", time.Duration(0), "duration value 1")
	fs.DurationVar(&d2, "d2", 1*time.Minute, "duration value 2")

	fs.Parse([]string{"-d1", "3s", "-d2", "1h2m3s"})

	output.Stdoutl("[d1]", *d1, (*d1).Milliseconds())
	output.Stdoutl("[d2]", d2, d2.Seconds())

	return nil

	/*
		$ task
		task: [build] go build .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: flags_duration

		[Name] "flags_duration"
		[d1]                 3s 3000
		[d2]                 1h2m3s 3723


		[Elapsed] 43.55µs
	*/

}
