package maps

import (
	"fmt"
	"strings"
)

// MapBasic -- マップについてのサンプル
func MapBasic() error {
	// マップの宣言
	map1 := make(map[int]string)
	map1[1] = "hoge"

	// キーの有無は以下のように判定する
	if v, ok := map1[1]; ok {
		fmt.Println(v)
	}

	// 関数もデータなので、値として設定可能
	map2 := make(map[string]func(string) string)
	map2["say"] = say
	map2["say2"] = say2

	if v, ok := map2["say"]; ok {
		fmt.Println(v("hoge"))
	}

	// マップをループする場合は以下のようにする
	for k, v := range map2 {
		var result string
		switch k {
		case "say":
			result = v("hoge")
		case "say2":
			result = v("hoge2")
		}

		fmt.Println(result)
	}

	return nil
}

func say(name string) string {
	return "hello " + strings.ToUpper(name)
}

func say2(name string) string {
	return "HELLO " + strings.ToUpper(name)
}
