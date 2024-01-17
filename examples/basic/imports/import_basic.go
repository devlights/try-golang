package imports

// 現状のコードでは利用していないが、後で利用するときなどはアンダーバーのエイリアスをつけておくと便利
// _ "fmt"

// Basic は、GO言語の import に関するサンプルです.
func Basic() error {
	// Go言語では「不要な宣言を許可しない」というポリシーがあるため
	// 利用していない import があると、コンパイルエラーになる
	// また、フォーマットを行うと利用していない import は削除される
	// しかし、現実ではデバッグの際などに一時的に特定の処理をコメントアウトに
	// して実行するなどが良くある。その度に本来必要である import が消されると
	// 面倒なときもある。
	//
	// そのような場合に GOでは
	//   - import するパッケージの左側に アンダーバー をつける
	// ことで、「参照されていないパッケージを強制的にプログラム内に組み込む」ことが
	// 出来るようになる。
	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: imports_basic

	   [Name] "imports_basic"


	   [Elapsed] 830ns
	*/

}
