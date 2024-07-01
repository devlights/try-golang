package flags

import (
	"flag"
	"net"
	"time"

	"github.com/devlights/gomy/output"
)

// TextVar は、 flag.TextVar() のサンプルです.
//
// flag.TextVar() は、Go 1.19 で追加された関数です。
// encoding.TextUnmarshaler を実装している型を指定出来ます。
//
// encoding.TextUnmarshaler を実装しているものとして以下のものがあります。
//
//   - time.Time
//   - net.IP
//   - url.URL
//
// # REFERENCES
//   - https://pkg.go.dev/flag@go1.22.4#TextVar
func TextVar() error {
	var (
		fs = flag.NewFlagSet("", flag.ExitOnError)
		tm time.Time
		ip net.IP
	)

	fs.TextVar(&tm, "time", time.Unix(0, 0), "time")
	fs.TextVar(&ip, "ip", net.IPv4(127, 0, 0, 1), "ip addr")

	// 時刻表記の末尾の Z はUTCを表す
	fs.Parse([]string{"-time", "2024-01-02T03:04:05Z", "-ip", "192.168.149.111"})

	tz, _ := time.LoadLocation("Asia/Tokyo")
	output.Stdoutf("[time]", "UTC=%s\tJST=%s\n", tm, tm.In(tz))
	output.Stdoutl("[ip  ]", ip)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: flags_textvar

	   [Name] "flags_textvar"
	   [time]               UTC=2024-01-02 03:04:05 +0000 UTC  JST=2024-01-02 12:04:05 +0900 JST
	   [ip  ]               192.168.149.111


	   [Elapsed] 152.56µs
	*/

}
