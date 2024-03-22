package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Println -- text/template の テンプレート仕様 における println 関数 のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Println() error {
	var (
		tmpls = []string{
			`{{ println . }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Println").Parse(t)
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

	   ENTER EXAMPLE NAME: templates_text_tmpl_println

	   [Name] "templates_text_tmpl_println"
	   [template]           {{ println . }}
	   [tmpl]               "100\n"
	   --------------------------------------------------


	   [Elapsed] 167.32µs
	*/

}
