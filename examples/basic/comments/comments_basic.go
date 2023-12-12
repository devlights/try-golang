/*
Package comments -- GO言語のコメントスタイルについてのサンプルがあるパッケージです

GO言語のコメントスタイルについては以下を参照.
  - http://bit.ly/2HS4sg4

スラッシュとアスタリスクを利用するコメントは「パッケージ用」

ダブルスラッシュを利用するコメントは「通常用、または、関数説明など」

- 全てのパッケージは、パッケージコメントを持つべきである。

(Every package should have a package comments)

- 関数などのドキュメントコメント(doc comments) は、そのアイテムの名前から始める。

(every doc comments begins with the name of the item it describes)
*/
package comments

import "fmt"

// Basic は、GO言語のコメントスタイルについてのサンプルです.
func Basic() error {
	// GO言語には、javadoc のようにソースコードのコメント
	// を利用してドキュメントを生成する機能がある。
	// 例えば以下のようにする
	//
	// $ go doc --all github.com/devlights/try-golang/basic/comments
	//
	// 上記のようにすると以下のように表示される
	goDocOutput := `
package comments // import "github.com/devlights/try-golang/basic/comments"

GO言語のコメントスタイルについてのサンプルがあるパッケージです

GO言語のコメントスタイルについては以下を参照.

    - https://tip.golang.org/doc/comment

スラッシュとアスタリスクを利用するコメントは「パッケージ用」

ダブルスラッシュを利用するコメントは「通常用、または、関数説明など」

- 全てのパッケージは、パッケージコメントを持つべきである。

(Every package should have a package comments)

- 関数などのドキュメントコメント(doc comments) は、そのアイテムの名前から始める。

(every doc comments begins with the name of the item it describes)

FUNCTIONS

func Basic() error
    Basic Basic は、GO言語のコメントスタイルについてのサンプルです.
`

	fmt.Print(goDocOutput)

	return nil

	/*
	    $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: comment_basic_usage

	   [Name] "comment_basic_usage"

	   package comments // import "github.com/devlights/try-golang/basic/comments"

	   GO言語のコメントスタイルについてのサンプルがあるパッケージです

	   GO言語のコメントスタイルについては以下を参照.

	       - http://bit.ly/2HS4sg4

	   スラッシュとアスタリスクを利用するコメントは「パッケージ用」

	   ダブルスラッシュを利用するコメントは「通常用、または、関数説明など」

	   - 全てのパッケージは、パッケージコメントを持つべきである。

	   (Every package should have a package comments)

	   - 関数などのドキュメントコメント(doc comments) は、そのアイテムの名前から始める。

	   (every doc comments begins with the name of the item it describes)

	   FUNCTIONS

	   func Basic() error
	       Basic Basic は、GO言語のコメントスタイルについてのサンプルです.


	   [Elapsed] 26.48µs
	*/

}
