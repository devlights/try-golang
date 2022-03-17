package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Call -- text/template の テンプレート仕様 における call のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Call() error {
	var (
		tmpls = []string{
			`{{ call .Fn 1 2 }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Call").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, struct{ Fn func(int, int) int }{func(x, y int) int { return x + y }})
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%q\n", buf.String())
		output.StdoutHr()
	}

	return nil
}
