// Package helloworld -- Go言語での Hello World プログラムが配置されているパッケージです。
package helloworld

import "fmt"

// Sync -- GO言語でのHelloWorldサンプル (同期版)
func Sync() error {
	// Golang には、 ビルドイン関数の println と
	// fmt.Println という　名前の似ている２つの関数があるが
	// 基本的に、どのサンプルも fmt.Println を利用している。
	//
	// 理由は、Golang のドキュメントに以下のように記載されているから。
	//   https://golang.org/builtin/#println
	//   https://qiita.com/taji-taji/items/79a49c0ee329d0b9c065
	for i := 0; i < 10; i++ {
		fmt.Printf("[%02d]\tHello World\n", i+1)
	}

	return nil
}
