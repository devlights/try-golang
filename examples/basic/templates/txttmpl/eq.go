package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Eq -- text/template の テンプレート仕様 における eq (equal) のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Eq() error {
	const (
		t1 = `{{$v := 10}} {{- if eq $v 10 -}} helloworld {{end}}`
		t2 = `{{$v := 10}} {{- if eq $v 11 -}} helloworld {{end}}`
		t3 = `{{$v := 10}} {{- if (eq $v 10) -}} helloworld {{end}}`
	)

	for _, t := range []string{t1, t2, t3} {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Eq").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, nil)
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%q\n", buf.String())
		output.StdoutHr()
	}

	return nil
}
