package command

import (
	"fmt"
	"path/filepath"
)

type (
	// Command は、サンプルの動作確認のためのインターフェースです.
	Command interface {
		Run() error
	}

	// ListFileCommand は、指定されたパターンのファイルを出力します.
	ListFileCommand struct {
		Dir     string
		Pattern string
	}
)

// インターフェースを実装できていることを検証するためのダミーフィールド
//
// 構造体が特定のインターフェースを実装できているかを検証し続けたい場合は
// 以下のようにEmptyなフィールドを作ってインターフェースの型で設定するようにしておく
// ポインタレシーバーの場合は, ゼロ値が nil なので以下のようになる.
// ポインタレシーバー出ない場合は、 ListFileCommand{} のように指定する
var _ Command = (*ListFileCommand)(nil)

/* 基本的に以下のようなインターフェースを返すコンストラクタを利用している場合はここで検証されるので心配ない
func NewListFileCommand(dir, pattern string) Command {
	c := new(listFileCommand)

	c.dir = dir
	c.pattern = pattern

	return c
}
*/

// Run は、パターンにマッチしたファイル名を出力します.
func (c *ListFileCommand) Run() error {
	matches, err := filepath.Glob(fmt.Sprintf("%s/%s", c.Dir, c.Pattern))
	if err != nil {
		return err
	}

	for _, v := range matches {
		fmt.Println(v)
	}

	return nil
}
