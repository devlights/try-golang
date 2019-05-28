package stdout

import "fmt"

// 標準出力についてのサンプル
func Println01() error {
	hello := "hello"
	world := "world"

	// fmt.Println() は、複数の値を指定するとスペースで区切って表示してくれる
	fmt.Println(hello, world)

	return nil
}
