package reflection

import (
	"reflect"

	"github.com/devlights/gomy/output"
)

// TypeOf -- reflect.TypeOf() のサンプル
func TypeOf() error {
	// 特定の値の型を調べる場合には、reflect.TypeOf() を利用する
	// C# でいう、 "hello world".GetType() と同じような感じ
	// Name メソッドで名前を取得できる
	i := 0
	t1 := reflect.TypeOf(i)
	output.Stdoutf("[int]", "%v\t%s\n", i, t1.Name())

	s := "hello world"
	t2 := reflect.TypeOf(s)
	output.Stdoutf("[string]", "%v\t%s\n", s, t2.Name())

	return nil
}
