package structtag

import (
	"reflect"
	"strings"

	"github.com/devlights/gomy/output"
)

// TagGet -- Struct Tag の内容を読み取るサンプルです。
//
// REFERENCES:
//   - https://text.baldanders.info/golang/struct-tag/
//   - https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go
//   - https://qiita.com/itkr/items/9b4e8d8c6d574137443c
//   - https://logmi.jp/tech/articles/324598
func TagGet() error {
	type data struct {
		val1 string `my:"upper"`
		val2 string `my:"lower"`
	}

	// Struct Tag の情報はリフレクション経由で取得する
	d := data{val1: "hello", val2: "WORLD"}
	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)

	// フィールドに定義されているタグ情報を取得して処理
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		// タグ情報取得
		tag := f.Tag.Get("my")

		// フィールドの値取得
		val := v.FieldByName(f.Name).String()

		// タグ情報に従った処理
		val2 := val
		switch tag {
		case "upper":
			val2 = strings.ToUpper(val)
		case "lower":
			val2 = strings.ToLower(val)
		}

		output.Stdoutf("[result]", "name: %v\ttag: %v\torig: %v\tconv: %v\n", f.Name, tag, val, val2)
	}

	return nil
}
