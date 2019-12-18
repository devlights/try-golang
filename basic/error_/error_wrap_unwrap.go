package error_

import (
	"errors"
	"fmt"
)

type (
	WrapError struct {
		InnerError error
	}
)

func (w *WrapError) Error() string {
	return fmt.Sprintf("example: WrapError(%v)", w.InnerError)
}

func (w *WrapError) Unwrap() error {
	return w.InnerError
}

func makeError(inner error) error {
	return fmt.Errorf("example: MakeError(%w)", inner)
}

func dump(e error) {
	fmt.Printf("[%T] %v\n", e, e)
}

// WrapAndUnwrap は、Goにおけるエラー処理にてエラーを内包するやり方についてのサンプルです。
// REFERENCES::
//   - https://medium.com/onefootball-locker-room/a-look-at-go-1-13-errors-9f6c9f6accb6
func WrapAndUnwrap() error {
	// ----------------------------------------------------------------
	// エラーの内包
	//
	// Goにて、エラーの内包（つまり、エラー情報がチェインしている状態）を実現する際は
	// 大抵、エラー型を定義して内部エラーを表すフィールドを用意して保持しておくパターンが多い。
	//
	// Go 1.13から、errors.Unwrap() の追加、及び、 %w 書式文字列の追加により
	// エラーの内包と内部エラーの取り出しについて共通手順が追加されたので、今後はこのやり方を
	// 使うべきである。
	// ----------------------------------------------------------------
	// %w を fmt.Errorf にて利用することで、簡単に埋め込みのエラーを作れる
	e1 := makeError(SentinelError)
	dump(e1)

	// 埋め込まれた内部エラーを取得
	e2 := errors.Unwrap(e1)
	dump(e2)

	// 独自に定義したエラー型でも同じやり方が出来る
	e3 := &WrapError{
		InnerError: &TypeAssertionError{},
	}
	dump(e3)

	// 埋め込まれた内部エラーを取得
	// 独自エラー型の場合は、Unwrap() error というシグネチャのメソッドを用意すると errors.Unwrap が呼んでくれる
	//
	// errors.Unwrap, errors.Is, errors.As は、同じ理屈で処理してくれるようになっており
	// それぞれ
	//   - interface { Unwrap() error }
	//   - interface { Is(error) bool }
	//   - interface { As(interface{}) bool }
	// を実装していれば、呼び出してくれる
	e4 := errors.Unwrap(e3)
	dump(e4)

	return nil
}
