package deepcopy

import (
	"encoding/json"
	"fmt"

	"github.com/devlights/gomy/output"
)

// JsonDeepCopy -- encoding/json を利用した deep-copy のサンプルです.
//
// REFERENCES:
//   - https://stackoverflow.com/questions/46790190/quicker-way-to-deepcopy-objects-in-golang
//   - https://stackoverflow.com/questions/37618399/efficient-go-serialization-of-struct-to-disk/37620399#37620399
//   - https://www.reddit.com/r/golang/comments/2vzmp7/deepcopy_of_a_slice_of_structs/
//   - https://groups.google.com/forum/#!topic/golang-nuts/vK6P0dmQI84
func JsonDeepCopy() error {
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
	// 以下は encoding/json を利用した方法
	// --------------------------------------------------------------------------
	pa := func(v interface{}) string {
		return fmt.Sprintf("%p", v)
	}

	clone := func(from, to interface{}) {
		b, _ := json.Marshal(from)
		_ = json.Unmarshal(b, to)
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

	output.Stdoutl("[i, i2]", i, i2, pa(&i), pa(&i2))
	output.Stdoutl("[s, s2]", s, s2, &s, &s2)

	// --------------------------------------------------------------------------
	// スライス
	var (
		sli1 = []int{1, 2, 3}
		sli2 []int
	)

	clone(&sli1, &sli2)
	output.Stdoutl("[sli1, sli2]", sli1, sli2, pa(&sli1), pa(&sli2))

	sli1[0] = 100
	output.Stdoutl("[sli1, sli2]", sli1, sli2, pa(&sli1), pa(&sli2))

	// --------------------------------------------------------------------------
	// マップ
	var (
		map1 = map[int]string{1: "apple", 2: "ringo"}
		map2 map[int]string
	)

	clone(&map1, &map2)
	output.Stdoutl("[map1, map2]", map1, map2, pa(&map1), pa(&map2))

	map1[1] = "melon"
	output.Stdoutl("[map1, map2]", map1, map2, pa(&map1), pa(&map2))

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
	output.Stdoutl("[b1, b2]", b1, b2, pa(&b1), pa(&b2))

	b1.Value = "world"
	b1.ValueB = "hello"
	output.Stdoutl("[b1, b2]", b1, b2, pa(&b1), pa(&b2))

	return nil
}
