package txttmpl

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/devlights/gomy/output"
)

// WhitespaceAndMinussign -- テンプレートで使用する {{- }} と {{ -}} についてのサンプルです。
//
//   - {{ -}} とすると、後続のスペースを除去する
//   - {{- }} とすると、直前のスペースを除去する
//
// template package の説明
//
//	テンプレートのソースコードを整形するために
//	アクションの左デリミタ（デフォルトでは"{{"）の直後にマイナス記号と空白が続く場合、直前のテキストからすべての末尾の空白が切り取られます。
//	同様に、右のデリミタ（"}}"）の前に空白とマイナス記号がある場合、直後のテキストからすべての先行する空白が切り取られます。
//
// REFERENCES
//   - https://pkg.go.dev/text/template@latest
func WhitespaceAndMinussign() error {
	const (
		t1 = `{{23}} < {{45}}`
		t2 = `{{23 -}} < {{45}}`
		t3 = `{{23}} < {{- 45}}`
		t4 = `{{23 -}} < {{- 45}}`
	)

	for i, t := range []string{t1, t2, t3, t4} {
		var (
			name = fmt.Sprintf("template-%d", i)
			tmpl = template.Must(template.New(name).Parse(t))
			buf  = bytes.Buffer{}
		)
		tmpl.Execute(&buf, nil)

		output.Stderrf("[tmpl]", "%s: %q\t%q\n", name, t, buf.String())
	}

	return nil
}
