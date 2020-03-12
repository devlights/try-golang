package effectivego

import "fmt"

// Semicolons -- Effective Go - Semicolons の 内容についてのサンプルです。
func Semicolons() error {
	/*
		https://golang.org/doc/effective_go.html#semicolons

		- Goでは、本来CやC#やJavaなどと同様にステートメントの最後にはセミコロンが必要
		  - だが、Goの構文解析器が自動でセミコロンを補ってくれているため、省略しても良い
		- このセミコロンの自動付与機能があるため、Goでは if や for などで C#のように {　を次の行に書くことが許されない
	*/

	// 以下はコンパイルエラーになる
	/*
		if true
		{
			fmt.Println("コンパイルエラー")
		}
	*/

	// 以下のように同じ行にしないといけない
	if true {
		fmt.Println("OK")
	}

	return nil
}
