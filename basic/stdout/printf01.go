package stdout

import "fmt"

// 標準出力についてのサンプル
func Printf01() error {
	message := "Hello Golang!!"
	fmt.Printf("%s", message)

	return nil
}
