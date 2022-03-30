package strconvs

import (
	"fmt"
	"strconv"
	"strings"
)

// HexToDecimal -- 16進数文字列を10進数に変換するサンプルです.
// (strconv.ParseInt() の 例)
func HexToDecimal() error {
	// -------------------------------------------------------------------
	// 16進数から10進数への変換
	//
	// Go にて、特定の進数文字列を変換するには
	// strconv.ParseInt() を使用する
	// -------------------------------------------------------------------
	hex := "0xffff"
	fmt.Printf("変換対象:[%s]\n", hex)

	// strconv.ParseInt() は、baseに16を指定した状態で
	// 16進数を変換する際に先頭に 0x が付与されていると失敗する
	_, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		fmt.Printf("変換失敗 %s\n", err)
	}

	// 先頭の 0x を除去
	hex = strings.Replace(hex, "0x", "", 1)

	i, err := strconv.ParseInt(hex, 16, 0)
	if err != nil {
		return err
	}
	fmt.Printf("変換結果 hex:[0x%v] -> dec:[%v]\n", hex, i)

	// baseに0を指定すると 0x の prefix が付与されていても成功する
	hex = "0xffff"
	i, err = strconv.ParseInt(hex, 0, 0)
	if err != nil {
		return err
	}
	fmt.Printf("変換結果 hex:[%v] -> dec:[%v]\n", hex, i)

	return nil
}
