package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Lt -- text/template の テンプレート仕様 における lt (less than) のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@go1.17.8
func Lt() error {
	const (
		// 10 < 10
		t1 = `{{ $v := 10 }} {{- if lt $v 10 -}} helloworld {{- end}}`
		// 10 < 11
		t2 = `{{ $v := 10 }} {{- if lt $v 11 -}} helloworld {{- end}}`
		// 10 < 9
		t3 = `{{ $v := 10 }} {{- if lt $v 9 -}} helloworld {{- end}}`
	)

	for _, t := range []string{t1, t2, t3} {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Lt").Parse(t)
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
