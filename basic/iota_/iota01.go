// REFERENCES:: http://bit.ly/2Re6Py5
package iota_

import "fmt"

type ByteSize int64

const (
	FirstValue int = iota // 最初の値は 0 からスタート
	SecondValue
	ThirdValue
)

const (
	_           = iota // 最初の iota の値は捨てる
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

func Iota01() error {
	fmt.Println(FirstValue, SecondValue, ThirdValue)
	fmt.Println(KB, MB, GB)
	return nil
}
