package error_

import (
	"errors"
	"fmt"
)

var (
	SentinelError = errors.New("example: sentinel error")
)

// Sentinel は、Goにおけるエラー処理イディオムの sentinel error check についてのサンプルです.
// REFERENCES::
//   - https://medium.com/onefootball-locker-room/a-look-at-go-1-13-errors-9f6c9f6accb6
func Sentinel() error {
	// ----------------------------------------------------------------
	// Sentinel error check
	//
	// Sentinelとは「衛兵」とか「ガード」の意味。
	// 予め、外部変数としてエラー変数を定義しておいて、発生したエラーをその変数の値と比べて
	// 一致している場合は、対象のエラーが発生していると見なすチェック方法.
	//
	// Go 1.13 から、 errors.Is() が追加されたので、従来では == を使って判定していた
	// 部分を errors.Is() で判定できるようになった。
	// ----------------------------------------------------------------
	raise := func() error {
		return SentinelError
	}

	if err := raise(); err != nil {

		// Go 1.13 より前のバージョンでは == で比較していた
		if err == SentinelError {
			fmt.Printf("Go 1.13 より前 (==で判定): %v\n", err)
		}

		// Go 1.13 からは errors.Is() があるので、そちらを使った方がシンプルかつ高機能
		if errors.Is(err, SentinelError) {
			fmt.Printf("Go 1.13 以降 (errors.Is()で判定): %v\n", err)
		}
	}

	return nil
}
