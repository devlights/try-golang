package stdout

import "fmt"

// Printf03 -- 標準出力についてのサンプル
func Printf03() error {
	data := MyData{30, "hello world"}

	fmt.Println(data.GetValue())
	fmt.Println(data.GetValue2())
	fmt.Println(data.GetValue3())
	fmt.Println(data.GetValue4())

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: printf03

	   [Name] "printf03"
	   &{30 hello world}
	   &{val1:30 val2:hello world}
	   &stdout.MyData{val1:30, val2:"hello world"}
	   *stdout.MyData


	   [Elapsed] 40.98µs
	*/

}

// MyData -- サンプル用の構造体
type MyData struct {
	val1 int
	val2 string
}

// GetValue -- 慣習として メソッドレシーバーの名前は構造体名の先頭一文字にする模様。
// self や this にすると警告が表示される
//
//	https://qiita.com/hnakamur/items/c99e3048f8902702a5a1
func (m *MyData) GetValue() string {
	return fmt.Sprintf("%v", m)
}

// GetValue2 -- 出力用メソッド
func (m *MyData) GetValue2() string {
	return fmt.Sprintf("%+v", m)
}

// GetValue3 -- 出力用メソッド
func (m *MyData) GetValue3() string {
	return fmt.Sprintf("%#v", m)
}

// GetValue4 -- 出力用メソッド
func (m *MyData) GetValue4() string {
	return fmt.Sprintf("%T", m)
}
