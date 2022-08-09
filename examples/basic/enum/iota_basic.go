package enum

import "fmt"

// 定数
const (
	FirstValue int = iota // 最初の値は 0 からスタート
	SecondValue
	ThirdValue
)

// ByteSize -- バイトサイズ
type ByteSize int64

// noinspection GoUnusedConst
const (
	_           = iota // 最初の iota の値は捨てる
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

// Weekday -- 曜日
type Weekday int

// 曜日
const (
	Sunday  Weekday = iota + 1 // iota は 0
	_                          // iota が加算
	Monday                     // iota は 2, (2 + 1) で 3 となる
	Tuesday                    // iota は 3, (3 + 1) で 4 となる
)

// Basic は、iota の基本的な使い方のサンプルです.
func Basic() error {
	fmt.Println(FirstValue, SecondValue, ThirdValue)
	fmt.Println(KB, MB, GB)
	fmt.Println(Sunday, Monday, Tuesday)
	return nil
}
