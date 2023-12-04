package errs

import (
	"errors"
	"fmt"
)

type (
	// TypeAssertionError -- type assertion check
	TypeAssertionError struct{}
)

func (t *TypeAssertionError) Error() string {
	return "example: type assertion error check"
}

// TypeAssertion は、Goにおけるエラー処理イディオムの type assertion check についてのサンプルです.
// REFERENCES::
//   - https://medium.com/onefootball-locker-room/a-look-at-go-1-13-errors-9f6c9f6accb6
func TypeAssertion() error {
	// ----------------------------------------------------------------
	// Type assertion check
	//
	// Sentinel error checkと似ているが、こちらは自前のエラー型を定義して
	// その方にキャストできるかどうかで、対象のエラーが発生したかどうかを判定するやり方。
	//
	// Go 1.13 から、 errors.As() が追加されたので、従来では 型検証 を使って判定していた
	// 部分を errors.As() で判定できるようになった。
	// ----------------------------------------------------------------
	raise := func() error {
		return &TypeAssertionError{}
	}

	if err := raise(); err != nil {

		// Go 1.13 より前のバージョンでは 型検証 で比較していた
		if _, ok := err.(*TypeAssertionError); ok {
			fmt.Printf("Go 1.13 より前 (型検証で判定): %v\n", err)
		}

		// Go 1.13 からは errors.As() があるので、そちらを使った方がシンプルかつ高機能
		// errors.As()は、少し使い方にクセがあるので注意。（2つ目の引数には、変換結果を入れる変数のポインタを渡す)
		var e *TypeAssertionError
		if errors.As(err, &e) {
			fmt.Printf("Go 1.13 以降 (errors.As()で判定): %v\n", e)
		}
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: error_typeassertion

	   [Name] "error_typeassertion"
	   Go 1.13 より前 (型検証で判定): example: type assertion error check
	   Go 1.13 以降 (errors.As()で判定): example: type assertion error check


	   [Elapsed] 9.089µs
	*/

}
