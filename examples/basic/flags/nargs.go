package flags

import (
	"flag"

	"github.com/devlights/gomy/output"
)

// Nargs は、flag.Arg(), flag.Args(), flag.NArg(), flag.NFlag() のサンプルです.
//
//   - flag.NFlag() は、処理したフラグの数
//   - flag.NArg()  は、処理していない引数の数
//   - flag.Arg(i)  は、処理していない引数のN番目を取得
//   - flag.Args()  は、処理していない引数リストを取得
//
// # REFERENCES
//   - https://pkg.go.dev/flag@go1.22.4#Arg
//   - https://pkg.go.dev/flag@go1.22.4#Args
//   - https://pkg.go.dev/flag@go1.22.4#NArg
//   - https://pkg.go.dev/flag@go1.22.4#NFlag
func Nargs() error {
	var (
		fs = flag.NewFlagSet("", flag.ContinueOnError)

		_ = fs.Int("i", 0, "int value")
		_ = fs.String("s", "", "string value")
	)

	fs.Parse([]string{"-i", "100", "-s", "hello", "out.txt", "out2.txt"})

	output.Stdoutl("[NFlag][処理したフラグの数    ]", fs.NFlag())
	output.Stdoutl("[NArg ][処理していない引数の数]", fs.NArg())

	for i := range fs.NArg() {
		output.Stdoutl("[Arg  ]", fs.Arg(i))
	}

	output.Stdoutl("[Args ]", fs.Args())

	return nil
}
