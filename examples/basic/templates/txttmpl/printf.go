package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Printf -- text/template の テンプレート仕様 における printf 関数 のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@go1.17.8
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
}
