package flags

import (
	"flag"

	"github.com/devlights/gomy/output"
)

// Visit は、flag.Visitのサンプルです。
//
//   - flag.Visit は、「実際にセットされたフラグを順に辞書順で走査してfnを呼ぶ」
//   - flag.VisitAll は、「セット有無に関わらずフラグを順に辞書順で走査してfnを呼ぶ」
//
// という動きになる。デフォルト値を持つフラグで「実際に指定された」かどうかを判定したい場合に使える。
//
// # REFERENCES
//   - https://pkg.go.dev/flag#Visit
func Visit() error {
	type (
		options struct {
			val1          int  // フラグとして使う値（指定される）
			val2          int  // フラグとして使う値（指定されない）
			val1Specified bool // 実際にval1フラグが指定されたかどうか
			val2Specified bool // 実際にval2フラグが指定されたかどうか
		}
	)
	var (
		opts options
		fs   = flag.NewFlagSet("", flag.ExitOnError)
	)
	fs.IntVar(&opts.val1, "v1", -1, "val1")
	fs.IntVar(&opts.val2, "v2", -1, "val2")

	var (
		args = []string{
			"-v1",
			"-1",
		}
	)
	fs.Parse(args)

	//
	// flag.Visit は、当然ながら flag.Parse してから呼び出さないと駄目
	// (parse前に呼び出す事もできるが、何もセットされていないので1回も呼ばれない)
	//
	var (
		fn = func(f *flag.Flag) {
			switch f.Name {
			case "v1":
				opts.val1Specified = true
			case "v2":
				opts.val2Specified = true
			}
		}
	)
	fs.Visit(fn)

	output.Stdoutl("[val1         ]", opts.val1)
	output.Stdoutl("[val1Specified]", opts.val1Specified)
	output.Stdoutl("[val2         ]", opts.val2)
	output.Stdoutl("[val2Specified]", opts.val2Specified)

	return nil

	/*
		$ task
		task: [build] go build -o "/home/dev/dev/github/try-golang/try-golang" .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: flags_visit

		[Name] "flags_visit"
		[val1         ]      -1
		[val1Specified]      true
		[val2         ]      -1
		[val2Specified]      false


		[Elapsed] 17.457µs
	*/
}
