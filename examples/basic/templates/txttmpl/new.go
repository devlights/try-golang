package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// New -- text/template の Newメソッドのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func New() error {
	//
	// テンプレートの使い方
	//
	// 1. template.New().Parse()
	// 2. *text.Template.Execute()
	//
	// Goのテンプレートでは {{ }} で囲っている部分が解釈される仕様となっている.
	// {{ . }} とすれば、*template.Template.Execute(io.Writer, interface{}) に指定した
	// データ (interface{}) そのものを指定していることになる。
	//
	// なので、 {{ .Message }} とかすると、指定したデータの Message フィールドの値を
	// 表示することが出来る
	//

	const (
		t = `hello {{.}}`
	)

	var (
		tmpl *template.Template
		buf  bytes.Buffer
		err  error
	)

	tmpl, err = template.New("New").Parse(t)
	if err != nil {
		return err
	}

	err = tmpl.Execute(&buf, "world")
	if err != nil {
		return err
	}

	output.Stdoutl("[template ]", t)
	output.Stdoutl("[tmpl-exec]", buf.String())

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_text_tmpl_new

	   [Name] "templates_text_tmpl_new"
	   [template ]          hello {{.}}
	   [tmpl-exec]          hello world


	   [Elapsed] 46.46µs
	*/

}
