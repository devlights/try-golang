package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// If -- text/template の テンプレート仕様 におけるifのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@go1.17.8
func If() error {
	const (
		t1 = `{{if .}} hello world {{end}}`
		t2 = `{{if . -}} hello world {{- end}}`
	)

	for _, t := range []string{t1, t2} {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("If").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, true)
		if err != nil {
			return err
		}
		output.Stdoutf("[true ]", "%q\n", buf.String())

		buf.Reset()
		err = tmpl.Execute(&buf, false)
		if err != nil {
			return err
		}
		output.Stdoutf("[false]", "%q\n", buf.String())
		output.StdoutHr()
	}

	return nil
}
