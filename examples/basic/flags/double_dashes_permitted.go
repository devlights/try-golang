package flags

import (
	"flag"

	"github.com/devlights/gomy/output"
)

// DoubleDashesPermitted は、flagパッケージにてデフォルトで "--flag" の指定が許容されていることを示すサンプルです。
//
// flagパッケージのドキュメントにて以下の形が許容されると記載があります。
//
//	-flag
//	--flag
//	-flag=x
//	-flag x
//
// # REFERENCES
//   - https://pkg.go.dev/flag@go1.25.0#hdr-Command_line_flag_syntax
func DoubleDashesPermitted() error {
	var (
		fs = flag.NewFlagSet("double-dashes", flag.ExitOnError)
		f1 = fs.Int("flag1", 0, "")
		f2 = fs.Int("flag2", 0, "")
		f3 = fs.Bool("flag3", false, "")
		f4 = fs.Bool("flag4", true, "")

		argv = []string{"-flag1", "10", "--flag2=20", "--flag3", "-flag4=false"}

		err error
	)
	if err = fs.Parse(argv); err != nil {
		return err
	}

	output.Stdoutl("[flag1]", *f1)
	output.Stdoutl("[flag2]", *f2)
	output.Stdoutl("[flag3]", *f3)
	output.Stdoutl("[flag4]", *f4)

	return nil

	/*
	   $ task

	   task: [build] go build -o "/workspace/try-golang/try-golang" .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: flags_double_dashes

	   [Name] "flags_double_dashes"
	   [flag1]              10
	   [flag2]              20
	   [flag3]              true
	   [flag4]              false

	   [Elapsed] 70.25µs
	*/
}
