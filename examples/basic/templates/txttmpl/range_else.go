package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// RangeElse -- text/template の テンプレート仕様 におけるrange .. elseのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@go1.17.8
//   - https://stackoverflow.com/questions/54156119/range-over-string-slice-in-golang-template
func RangeElse() error {
	const (
		t1 = `{{range .Items}} {{.}} {{end}}`
		t2 = `{{range .Items}} {{.}} {{else}} no item {{end}}`
	)
	type (
		D struct {
			Items []string
		}
	)

	for _, t := range []string{t1, t2} {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Range").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, D{[]string{}})
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%q\n", buf.String())
		output.StdoutHr()
	}

	return nil
}
