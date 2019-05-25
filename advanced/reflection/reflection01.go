// reflect パッケージのサンプルです。
//
// 型オブジェクトの取得について。
// 参考情報：http://bit.ly/2UON9BD
package reflection

import (
	"fmt"
	"reflect"
)

// reflect.TypeOf() のサンプルです。
func Reflection01() error {
	// 特定の値の型を調べる場合には、reflect.TypeOf() を利用する
	// C# でいう、 "hello world".GetType() と同じような感じ
	// Name メソッドで名前を取得できる
	i := 0
	t1 := reflect.TypeOf(i)
	fmt.Printf("%s\n", t1.Name())

	s := "hello world"
	t2 := reflect.TypeOf(s)
	fmt.Printf("%s\n", t2.Name())

	return nil
}
