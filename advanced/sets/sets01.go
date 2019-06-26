package sets

import (
	"fmt"
	"github.com/deckarep/golang-set"
)

// Set01 - mapset.NewSet() の サンプル
// REFERENCES:: https://github.com/deckarep/golang-set
func Set01() error {

	// 新しい sets を生成
	s1 := mapset.NewSet()

	// データを設定
	s1.Add("hello")
	s1.Add("world")
	s1.Add("hello")

	// 集合なので重複項目は存在しない
	fmt.Println(s1)

	return nil
}
