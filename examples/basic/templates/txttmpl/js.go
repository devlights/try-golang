package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Js -- text/template の テンプレート仕様 における js のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Js() error {
	var (
		tmpls = []string{
			`{{ printf "%s:%s" . (js .) }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Js").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, struct{ Hello, World string }{"><://=&", "world"})
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

	   ENTER EXAMPLE NAME: templates_text_tmpl_js

	   [Name] "templates_text_tmpl_js"
	   [template]           {{ printf "%s:%s" . (js .) }}
	   [tmpl]               "{><://=& world}:{\\u003E\\u003C://\\u003D\\u0026 world}"
	   --------------------------------------------------


	   [Elapsed] 114.32µs
	*/

}
