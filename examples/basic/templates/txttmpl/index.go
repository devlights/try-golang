package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Index -- text/template の テンプレート仕様 における index のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Index() error {
	var (
		tmpls = []string{
			`{{ printf "%v\t[1][2]=%d" . (index . 1 2) }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Index").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, [2][3]int{{1, 2, 3}, {3, 4, 5}})
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%s\n", buf.String())
		output.StdoutHr()
	}

	return nil
}
