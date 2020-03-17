package effectivego15

import (
	"github.com/devlights/gomy/output"
)

// Maps -- Effective Go - Maps の 内容についてのサンプルです。
func Maps() error {
	/*
		https://golang.org/doc/effective_go.html#maps

		- Goではマップは組み込み型として用意されている
		- 使い方は他の言語と同じ
		- スライスはキーとして使えない
		- スライスと同様にマップも内部に保持しているデータへの参照を持っている
	*/
	// マップの宣言
	// 上から順に make(), 中身を空で生成, 初期値を設定して生成
	m1 := make(map[string]int)
	m2 := map[string]int{}
	m3 := map[string]int{
		"Go":     1,
		"C#":     2,
		"Python": 3,
	}

	output.Stdoutl("(1)", m1, m2, m3)

	// マップは内部にデータへの参照を保持しているので
	// 単純に代入すると、当然同じ参照をみることになる.
	languages := m3
	languages["Go"] = 999

	output.Stdoutl("(2)", m3, languages)

	// マップを deep copy するには、単純にループする
	// （スライスの場合は 組み込みの copy() があるので、それを使う)
	languages2 := map[string]int{}
	for key, value := range languages {
		languages2[key] = value
	}

	languages["Go"] = 888

	output.Stdoutl("(3)", m3, languages, languages2)

	// 存在しないキーにアクセスすると ゼロ値が返る
	// (例： int の場合は0, string の場合は"")
	output.Stdoutl("(4)", languages["not_exists"])

	// マップの値取得の際は、value, bool の２つを受け取れる
	// 左辺の値が一つの場合は value のみ。
	// ２つ指定している場合は、value, bool が返る
	// ２つ目の bool の値が true の場合、キーが存在している
	// ２つ目の 変数名 は、暗黙でokという名前にすることが多い
	_, ok := languages["not_exists"]
	if !ok {
		output.Stdoutl("(5)", "key [not_exists] is not exists.")
	}

	// 上のイディオムはGoではまとめて以下のようにすることが多い
	if _, ok = languages["not_exists"]; !ok {
		output.Stdoutl("(6)", "key [not_exists] is not exists.")
	}

	// マップのエントリを削除する場合は
	// 組み込み関数の delete() を使用する
	delete(languages, "Python")

	output.Stdoutl("(7)", m3, languages, languages2)

	return nil
}
