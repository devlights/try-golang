package maps

import "fmt"

func MapAccess() error {

	var (
		f = func(key string, val int, ok bool) {
			fmt.Printf("m[%v]\tval: %v\tok: %v\n", key, val, ok)
		}

		m = map[string]int{
			"a": 100,
			"b": 200,
		}
	)

	// -------------------------------------------------------------
	// Go の場合、mapに対応するキーが存在するかどうかの確認は
	// 実際に map に対してアクセスした際に戻り値で返ってくる bool の値で判別できる
	// -------------------------------------------------------------
	// 存在する場合
	val, ok := m["a"]
	f("a", val, ok)

	// 存在しない場合
	val, ok = m["not_exists"]
	f("not_exists", val, ok)

	// このパターンには定型が存在する. この書き方はGoでよく見る書き方である.
	if val, ok = m["a"]; ok {
		f("存在するパターン", val, ok)
	}

	if val, ok = m["not_exists"]; !ok {
		f("存在しないパターン", val, ok)
	}

	return nil
}
