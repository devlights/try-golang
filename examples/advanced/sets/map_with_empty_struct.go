package sets

import "fmt"

func MapSet01() error {

	// map を set の代用として利用できる
	// Keyを set で利用する値
	// Valueを 空の構造体 にする
	//
	// golang にて 空の構造体 (struct{}) は
	// 0バイトになる.
	//
	// REFERENCES:: http://bit.ly/2ZL7A4N
	// 				http://bit.ly/2ZL8yhr

	// 例えば、数値の集合をつくるとすると以下のようにする
	set := make(map[int]struct{})

	// 値の設定はちょっと変な感じだかこんな風になる
	// struct{} で 一つの型。interface{} と同じ。
	// なので、初期化するために struct{}{} とする。
	set[1] = struct{}{}
	set[2] = struct{}{}
	set[1] = struct{}{}

	fmt.Println(set)

	// ループする場合、 key のみ取得で良い
	for k := range set {
		fmt.Println(k)
	}

	return nil
}
