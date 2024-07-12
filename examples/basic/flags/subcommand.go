package flags

import (
	"flag"
	"fmt"
)

// Subcommand は、flagパッケージを使ってサブコマンドを実現するサンプルです.
//
// flagパッケージを使って、サブコマンドを実現する場合は
// flag.FlagSetでそれぞれのサブコマンドを表現する。
//
// サブコマンドの数が多い場合や細かな制御などをしたい場合は
// [flaggy](https://github.com/integrii/flaggy) などの外部ライブラリを利用する方が楽。
//
// flaggyのサンプルについては [try-golang-extlib](https://github.com/devlights/try-golang-extlib/tree/main/examples/singleapp/flaggy) を参照。
//
// # REFERENCES
//   - https://oohira.github.io/gobyexample-jp/command-line-subcommands.html
//   - https://pkg.go.dev/flag
//   - https://github.com/integrii/flaggy
func Subcommand() error {
	var (
		sub1 = flag.NewFlagSet("sub1", flag.ExitOnError) // サブコマンド
		sub2 = flag.NewFlagSet("sub2", flag.ExitOnError) // サブコマンド

		sub1BoolVal bool
		sub2IntVal  int
	)
	sub1.BoolVar(&sub1BoolVal, "enable", false, "bool flag")
	sub2.IntVar(&sub2IntVal, "value", 0, "int flag")

	//
	// 引数の１つ目がサブコマンドとなるので、それで振り分け
	//
	var (
		argsList = [][]string{
			{"sub1"},
			{"sub2"},
			{"sub1", "-enable"},
			{"sub2", "-value", "100"},
		}
	)
	for _, args := range argsList {
		switch args[0] {
		case "sub1":
			sub1.Parse(args[1:])
			fmt.Printf("[sub1] enable=%v\n", sub1BoolVal)
		case "sub2":
			sub2.Parse(args[1:])
			fmt.Printf("[sub2] value=%v\n", sub2IntVal)
		}
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: flags_subcommand

	   [Name] "flags_subcommand"
	   [sub1] enable=false
	   [sub2] value=0
	   [sub1] enable=true
	   [sub2] value=100


	   [Elapsed] 37.48µs
	*/

}
