package txttmpl

import (
	"fmt"
	"os"
	"text/template"
)

// MissingKey -- text/template の missingkey=zero オプション指定時のサンプルです。
//
// Template.Option("missingkey=zero")の設定を行い、テンプレートに対して map[string]any をデータとして指定した場合
// キーが存在しないの場合、"<no value>" と表示されることに注意。（マップの値の方の方がanyのため）
//
// 当然、any以外の場合はその型のゼロ値となる。
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func MissingKey() error {
	const (
		TEXT = `{{.Value}} world`
	)
	var (
		mAny = map[string]any{"Value": "hello"}
		mStr = map[string]string{"Value": "hello"}
		mInt = map[string]int{"Value": 100}

		t  = template.Must(template.New("MissingKey").Option("missingkey=zero").Parse(TEXT))
		fn = func(message string) {
			fmt.Printf("[%s] -----------------------\n", message)
			t.Execute(os.Stdout, mAny)
			fmt.Println("")
			t.Execute(os.Stdout, mStr)
			fmt.Println("")
			t.Execute(os.Stdout, mInt)
			fmt.Println("")
		}
	)
	fn("キーあり")

	delete(mAny, "Value")
	delete(mStr, "Value")
	delete(mInt, "Value")

	fn("キーなし")

	return nil

	/*
	   $ task
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_text_tmpl_missingkey

	   [Name] "templates_text_tmpl_missingkey"
	   [キーあり] -----------------------
	   hello world
	   hello world
	   100 world
	   [キーなし] -----------------------
	   <no value> world
	   	world
	   0 world

	   [Elapsed] 79.082µs
	*/
}
