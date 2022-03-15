package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// ElseIf -- text/template の テンプレート仕様 におけるelseifのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@go1.17.8
//   - https://ema-hiro.hatenablog.com/entry/20170729/1501320887
//   - https://stackoverflow.com/questions/33282061/go-templates-with-eq-and-index
//   - https://stackoverflow.com/questions/31101729/compare-strings-in-templates
func ElseIf() error {
	const (
		t1 = `{{if eq .Message "hello" -}} world {{- else if eq .Message "world" -}} hello {{- end}}`
	)
	type (
		D struct {
			Message string
		}
	)

	for _, t := range []string{t1} {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("ElseIf").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, D{"world"})
		if err != nil {
			return err
		}

		output.Stdoutl("[tmpl]", buf.String())
		output.StdoutHr()
	}

	return nil
}
