package stdout

import "fmt"

// 標準出力についてのサンプル
func Printf03() error {
	data := MyData{30, "hello world"}

	fmt.Println(data.GetValue())
	fmt.Println(data.GetValue2())
	fmt.Println(data.GetValue3())
	fmt.Println(data.GetValue4())

	return nil
}

// サンプル用の構造体
type MyData struct {
	val1 int
	val2 string
}

// 慣習として メソッドレシーバーの名前は構造体名の先頭一文字にする模様。
// self や this にすると警告が表示される
//     https://qiita.com/hnakamur/items/c99e3048f8902702a5a1
func (m *MyData) GetValue() string {
	return fmt.Sprintf("%v", m)
}

func (m *MyData) GetValue2() string {
	return fmt.Sprintf("%+v", m)
}

func (m *MyData) GetValue3() string {
	return fmt.Sprintf("%#v", m)
}

func (m *MyData) GetValue4() string {
	return fmt.Sprintf("%T", m)
}
