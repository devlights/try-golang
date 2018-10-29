// HelloWorld プログラム
package main

import "fmt"

func main() {
	// Golang には、 ビルドイン関数の println と
	// fmt.Println という　名前の似ている２つの関数があるが
	// 基本的に、どのサンプルも fmt.Println を利用している。
	//
	// 理由は、Golang のドキュメントに以下のように記載されているから。
	//   https://golang.org/pkg/builtin/#println
	//   https://qiita.com/taji-taji/items/79a49c0ee329d0b9c065
	fmt.Println("Hello World!")
}
