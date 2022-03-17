package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Newline -- text/template の テンプレート仕様 における改行のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
//   - https://zenn.dev/moutend/articles/5a7b6f6e0c185fde1716
//   - https://stackoverflow.com/questions/39948383/how-to-avoid-newlines-caused-by-conditionals
func Newline() error {
	const (
		t1 = `hello {{ printf "\n" }} world`
		t2 = `hello {{- printf "\n" -}} world`
	)

	for _, t := range []string{t1, t2} {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Newline").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, nil)
		if err != nil {
			return err
		}

		output.Stdoutf("[%q]", "%q\n", buf.String())
		output.Stdoutf("[%s]", "%s\n", buf.String())
		output.StdoutHr()
	}

	return nil
}
