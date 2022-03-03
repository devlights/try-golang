package reflects

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/devlights/gomy/output"
)

// FindType -- 実行時に型を求めるやり方についてのサンプルです
//
// REFERENCES:
//   - https://stackoverflow.com/questions/20170275/how-to-find-the-type-of-an-object-in-go
//   - https://pkg.go.dev/reflect
func FindType() error {
	type (
		MyType struct {
			Value int
		}
	)

	var (
		v = MyType{100}
	)

	//
	// 1. 文字列フォーマットの機能を利用
	//
	t1 := fmt.Sprintf("%T", v)
	output.Stdoutl("[string formatting]", t1)

	//
	// 2. reflect.TypeOf を利用
	//
	t2 := reflect.TypeOf(v).String()
	output.Stdoutl("[reflect typeof]", t2)

	//
	// 3. type assertions を利用
	//
	var (
		obj interface{} = v
	)

	switch obj.(type) {
	case MyType:
		output.Stdoutl("[type assertsion]", "reflects.MyType")
	default:
		return errors.New("型がおかしい")
	}

	return nil
}
