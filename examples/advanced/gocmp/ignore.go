package gocmp

import (
	"reflect"
	"time"

	"github.com/devlights/gomy/output"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type (
	// D -- サンプル用構造体
	D struct {
		StrField     string    // パブリックな文字列フィールド
		TimeField    time.Time // パブリックな日付フィールド
		privateField string    // プライベートなフィールド
	}
)

// Ignore -- go-cmp にて 指定したフィールド を無視して比較するサンプルです.
func Ignore() error {
	// ---------------------------------------------------------------------------------
	// relect.DeepEqual() は便利であるが、以下の問題がある
	//   - unexported な フィールド まで比較してしまう
	//   - time.Time な フィールド があると一致判定ができない
	//
	// go-cmp では unexported な フィールドを無視したり
	// 明示的に特定のフィールドを無視したり出来る
	// ---------------------------------------------------------------------------------
	var (
		now = time.Now()
		d1  = D{
			StrField:     "hello",
			TimeField:    now,
			privateField: "world",
		}
		d2 = D{
			StrField:     "hello",
			TimeField:    now.Add(2 * time.Second),
			privateField: "golang",
		}
	)

	// ---------------------------------------------------------------------------------
	// 上の d1 と d2 は以下の状態
	//
	// - プライベートなフィールドの値は異なる
	// - StrField は 同じ値
	// - TimeField は 異なる値
	//
	// このようなオブジェクトの場合、日付の値は比較対象から除外して
	// さらにプライベートな値も比較対象から除外して、意味のある公開プロパティが
	// 一致している場合は同値のオブジェクトであると判定することがよくある。
	//
	// ---------------------------------------------------------------------------------

	// ---------------------------------------------------------------------------------
	// reflect.DeepEqual で 比較すると 当然 false となる
	output.Stdoutl("reflect.DeepEqual(d1, d2)", reflect.DeepEqual(d1, d2))

	// ---------------------------------------------------------------------------------
	// go-cmp にて以下の設定を施し比較する
	//
	// - TimeField は 日付情報 なので無視するフィールドとする
	// - プライベートなフィールドは比較しない
	//
	opts := cmp.Options{
		cmpopts.IgnoreFields(d1, "TimeField"),
		cmpopts.IgnoreUnexported(d1),
	}

	output.Stdoutl("cmp.Equal(d1, d2, opts)", cmp.Equal(d1, d2, opts))

	return nil
}
