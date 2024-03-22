package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Printf -- text/template の テンプレート仕様 における printf 関数 のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Printf() error {
	var (
		tmpls = []string{
			`{{ printf "%T:%v" . . }}`,
			`{{ printf "%[1]T:%[1]v" . }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Printf").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, 100)
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%q\n", buf.String())
		output.StdoutHr()
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_text_tmpl_printf

	   [Name] "templates_text_tmpl_printf"
	   [template]           {{ printf "%T:%v" . . }}
	   [tmpl]               "int:100"
	   --------------------------------------------------
	   [template]           {{ printf "%[1]T:%[1]v" . }}
	   [tmpl]               "int:100"
	   --------------------------------------------------


	   [Elapsed] 169.94µs
	*/

}
