package txttmpl

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/devlights/gomy/output"
)

// FuncMap -- text/template の テンプレート仕様 における FuncMap (独自関数の登録) のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
//   - https://stackoverflow.com/questions/38081807/for-loop-of-two-variables-in-go
func FuncMap() error {
	var (
		tmpls = []string{
			`{{ upper . }}  {{ upper (fn .) }}`,
		}
		fn = func(v interface{}) string {
			s, ok := v.(string)
			if !ok {
				return ""
			}

			var (
				result  = make([]byte, len(s))
				lastIdx = len(s) - 1
			)

			for i, j := lastIdx, 0; i >= 0; i, j = i-1, j+1 {
				result[j] = s[i]
			}

			return string(result)
		}
		funcs = template.FuncMap{
			"fn":    fn,
			"upper": strings.ToUpper,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		// 独自関数を利用する場合は Parse の呼び出しの前に Funcs を呼び出す必要がある
		tmpl, err = template.New("FuncMap").Funcs(funcs).Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, "helloworld")
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%s\n", buf.String())
		output.StdoutHr()
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_text_tmpl_funcmap

	   [Name] "templates_text_tmpl_funcmap"
	   [template]           {{ upper . }}  {{ upper (fn .) }}
	   [tmpl]               HELLOWORLD  DLROWOLLEH
	   --------------------------------------------------


	   [Elapsed] 89.18µs
	*/

}
