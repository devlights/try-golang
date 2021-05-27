package shortassignment

import (
	"fmt"
	"os"
	"path/filepath"
)

// Basic -- ":=" を使用した簡易変数初期化方法についてのサンプル
func Basic() error {
	// 関数の内部でだけ、 ":=" を用いて暗黙変数を初期化することができる
	// これが利用できるのは関数の内側だけで、関数の外側では すべての文 はキーワードで始まらないと駄目
	// Goを使っていると頻繁に利用するものである。他の言語でいう var に近い。
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	f := filepath.Join(dir, "basic", "variables", "short_assignment_statement.go")

	finfo, err := os.Stat(f)
	if err != nil {
		return err
	}

	fmt.Printf("現在の作業ディレクトリ: %s\nファイルパス: %s\nサイズ: %d bytes\n", dir, f, finfo.Size())

	return nil
}
