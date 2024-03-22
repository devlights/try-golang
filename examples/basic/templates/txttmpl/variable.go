package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Variable -- text/template の テンプレート仕様 における 変数 のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Variable() error {
	var (
		tmpls = []string{
			`{{ $v := . }}{{ $i := 10 }}{{ printf "%v\t%d" $v $i }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Variable").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, []int{1, 2, 3})
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%s\n", buf.String())
		output.StdoutHr()
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_text_tmpl_variable

	   [Name] "templates_text_tmpl_variable"
	   [template]           {{ $v := . }}{{ $i := 10 }}{{ printf "%v\t%d" $v $i }}
	   [tmpl]               [1 2 3]    10
	   --------------------------------------------------


	   [Elapsed] 128.38µs
	*/

}
