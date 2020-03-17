package effectivego02

// Formatting -- Effective Go - Formatting の 内容についてのサンプルです。
func Formatting() error {
	/*
		https://golang.org/doc/effective_go.html#formatting

		- Go ではソースコードのフォーマットの問題に対してツールで対応する方針を持っている
			- go fmt
			- gofmt
		- go fmt がパッケージ向け。 go fmt は内部で gofmt を呼び出している。
		- Go ではインデントにタブを使う
		- Go では一行の文字数を規定していない。 PythonのPEP8とかだと規定がある。
		- なので、 Go で作業している際はコメントの位置合わせなどを気にする必要はない
		- ソースコードを書いたら go fmt するようにする
			- IDE (Golandなど) や VSCode などを利用していると自動でやってくれる
	*/

	return nil
}
