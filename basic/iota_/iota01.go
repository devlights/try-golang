// REFERENCES:: http://bit.ly/2Re6Py5
package iota_

import "fmt"

const (
	FirstValue int = iota // 最初の値は 0 からスタート
	SecondValue
	ThirdValue
)

type ByteSize int64

//noinspection GoUnusedConst
const (
	_           = iota // 最初の iota の値は捨てる
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

type Weekday int

const (
	Sunday  Weekday = iota + 1 // iota は 0
	_                          // iota が加算
	Monday                     // iota は 2, (2 + 1) で 3 となる
	Tuesday                    // iota は 3, (3 + 1) で 4 となる
)

func Iota01() error {
	fmt.Println(FirstValue, SecondValue, ThirdValue)
	fmt.Println(KB, MB, GB)
	fmt.Println(Sunday, Monday, Tuesday)
	return nil
}
