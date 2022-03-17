package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Ne -- text/template の テンプレート仕様 における ne (not equal) のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Ne() error {
	const (
		t1 = `{{$v := 10}} {{- if ne $v 10 -}} helloworld {{end}}`
		t2 = `{{$v := 10}} {{- if ne $v 11 -}} helloworld {{end}}`
		t3 = `{{$v := 10}} {{- if (ne $v 10) -}} helloworld {{end}}`
	)

	for _, t := range []string{t1, t2, t3} {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Ne").Parse(t)
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
