package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Slice -- text/template の テンプレート仕様 における slice のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Slice() error {
	var (
		tmpls = []string{
			`{{ print (slice .) }}`,
			`{{ print (slice . 1) }}`,
			`{{ print (slice . 1 3) }}`,
			`{{ print (slice . 1 2 2) }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Slice").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, []string{"hello", "go", "golang", "world"})
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%s\n", buf.String())
		output.StdoutHr()
	}

	return nil
}
