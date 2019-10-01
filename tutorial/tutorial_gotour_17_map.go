package tutorial

import "fmt"

// Map は、 Tour of Go - Maps (https://tour.golang.org/moretypes/19) の サンプルです。
func Map() error {
	// ------------------------------------------------------------
	// Go言語のマップ
	// Go言語のマップは、他の言語のマップ/ディクショナリと同じもの.
	// Go言語のマップの型指定は、少し特殊で
	//   map[キーの型]値の型
	// という形で宣言する.
	//
	// マップのゼロ値は スライス と同様に nil
	// マップを生成する場合は、初期化リテラルを利用するか組み込み関数 make() を使う
	// キーが存在するかどうかの確認は、すこし特殊な判定の仕方を行う
	//   if v, ok := map1["not_exists"]; ok {
	//     // キーが存在する
	//   }
	// マップから要素を削除する場合は、組み込み関数 delete() を使用する
	//
	// ------------------------------------------------------------
	var (
		map1 map[int]string
	)

	// マップの初期値は nil
	fmt.Printf("%#v\n", map1)

	// マップを生成
	map1 = make(map[int]string)
	fmt.Printf("%#v\n", map1)

	// 値を設定
	map1[100] = "apple"
	fmt.Printf("%#v\n", map1)

	// 値を更新
	map1[100] = "banana"
	fmt.Printf("%#v\n", map1)

	// キーが存在するか確認
	if v, ok := map1[100]; ok {
		fmt.Println(v)
	}

	// 値の削除
	delete(map1, 100)
	fmt.Println(map1)

	// マップのリテラル表記
	var (
		map2 = map[int]string{
			100: "apple",
		}
	)

	fmt.Println(map2)

	return nil
}
