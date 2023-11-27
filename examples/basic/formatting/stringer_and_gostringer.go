package formatting

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type _MyInt int

func (me _MyInt) String() string {
	return fmt.Sprintf("stringer: %d", me)
}

func (me _MyInt) GoString() string {
	return fmt.Sprintf("gostring: %d", me)
}

// StringerAndGoStringer は、fmt.Stringerとfmt.GoStringerについてのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/fmt@go1.21.4#Stringer
//   - https://pkg.go.dev/fmt@go1.21.4#GoStringer
func StringerAndGoStringer() error {
	//
	// fmt.Stringer と fmt.GoStringer の違い
	//
	// どちらも文字列表現のためのインターフェースであるが
	// fmt.GoStringerの方は GoString() を呼ばれた場合に発動する.
	// これは、fmt.Printf などで %#v を利用してフォーマットしようとする際に使われる
	//

	var (
		i    = _MyInt(100)
		s    = strconv.Itoa((int)(i))
		b, _ = json.Marshal(i)
	)

	fmt.Printf("v: %v\t+v: %+v\t#v: %#v\tjson: %s\ts: %s", i, i, i, b, s)

	return nil
}
