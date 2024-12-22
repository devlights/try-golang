package deepcopy

import (
	"bytes"
	"encoding/gob"

	"github.com/devlights/gomy/output"
)

// GobDeepCopy -- encoding/gob を利用した deep-copy のサンプルです.
//
// REFERENCES:
//   - https://stackoverflow.com/questions/46790190/quicker-way-to-deepcopy-objects-in-golang
//   - https://stackoverflow.com/questions/37618399/efficient-go-serialization-of-struct-to-disk/37620399#37620399
//   - https://www.reddit.com/r/golang/comments/2vzmp7/deepcopy_of_a_slice_of_structs/
//   - https://groups.google.com/forum/#!topic/golang-nuts/vK6P0dmQI84
func GobDeepCopy() error {
	// --------------------------------------------------------------------------
	// encoding/gob パッケージを利用した deep-copy
	//
	// 他の言語でもよくやる方法であるが、手っ取り早くオブジェクトを deep-copy する方法として
	// 対象のオブジェクトを シリアライズ/デシリアライズ して取得する方法がある。
	//
	// golang にて シリアライズ/デシリアライズ するメジャーな方法が以下のもの
	//   - encoding/gob 使う
	//   - encoding/json 使う
	//
	// 以下は encoding/gob を利用した方法
	//
	// gobは、Go言語で表現される型のほどんどに対応しているが、以下の型は対応していない。
	//	 - 関数型 (func)
	//	 - チャネル型（chan）
	//	 - エクスポートされていないフィールド
	// 上記の型の場合、定義されていても無視される。
	//
	// また、gobは最初の出力時に型の情報を出力するようになっているため
	// 単発でディープコピーをする場合には encoding/json の方が速い。(jsonにはこのフェーズが無いため)
	// --------------------------------------------------------------------------
	clone := func(from, to any) {
		var (
			buf = new(bytes.Buffer)
			enc = gob.NewEncoder(buf)
			dec = gob.NewDecoder(buf)
		)
		_ = enc.Encode(from)
		_ = dec.Decode(to)
	}

	// --------------------------------------------------------------------------
	// 基本の型
	var (
		i  = 100
		i2 int

		s  = "helloworld"
		s2 string
	)

	clone(&i, &i2)
	clone(&s, &s2)

	output.Stdoutl("[i, i2]", i, i2)
	output.Stdoutl("[s, s2]", s, s2)

	// --------------------------------------------------------------------------
	// スライス
	var (
		sli1 = []int{1, 2, 3}
		sli2 []int
	)

	clone(&sli1, &sli2)
	output.Stdoutl("[sli1, sli2][1]", sli1, sli2)

	sli1[0] = 100
	sli2[1] = 200
	output.Stdoutl("[sli1, sli2][2]", sli1, sli2)

	// --------------------------------------------------------------------------
	// マップ
	var (
		map1 = map[int]string{1: "apple", 2: "ringo"}
		map2 map[int]string
	)

	clone(&map1, &map2)
	output.Stdoutl("[map1, map2][1]", map1, map2)

	map1[1] = "melon"
	map2[2] = "林檎"
	output.Stdoutl("[map1, map2][2]", map1, map2)

	// --------------------------------------------------------------------------
	// 構造体
	type (
		A struct {
			Value string
		}

		B struct {
			A
			ValueB string
		}
	)

	var (
		b1 = B{
			A:      A{Value: "hello"},
			ValueB: "world",
		}
		b2 B
	)

	clone(&b1, &b2)
	output.Stdoutl("[b1, b2][1]", b1, b2)

	b1.Value = "world"
	b1.ValueB = "hello"
	output.Stdoutl("[b1, b2][2]", b1, b2)

	return nil
}
