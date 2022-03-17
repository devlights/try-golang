package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Must -- Template.Must() のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Must() error {
	const (
		t = `hello {{.}}`
	)

	var (
		tmpl *template.Template
		buf  bytes.Buffer
		err  error
	)

	//
	// template.Must() は、引数に template.New().Parse() の結果を
	// そのまま渡すことが出来るように設計されているヘルパー関数。
	//
	// テンプレートを初期化する際にエラーが発生することは本番時には
	// ほぼありえない。なので、このヘルパー関数を通すことで err の
	// 処理をしなくて済む。
	//
	// もし、Must関数を使った際にエラーが発生した場合は
	// panic するので要注意。
	//
	tmpl = template.Must(template.New("Must").Parse(t))

	err = tmpl.Execute(&buf, "world")
	if err != nil {
		return err
	}

	output.Stdoutl("[template ]", t)
	output.Stdoutl("[tmpl-exec]", buf.String())

	return nil
}
