package gocmp

import (
	"strings"

	"github.com/devlights/gomy/output"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type (
	// A -- サンプル用構造体
	A struct {
		PublicField string
	}

	// B -- サンプル用構造体
	B struct {
		A
		privateField string
	}

	// C -- サンプル用構造体
	C A
)

// Equal -- 自身と指定された値が等しいかどうかを返します
//
// 等しい場合は true, それ以外は false です.
func (c C) Equal(o C) bool {
	upper1 := strings.ToUpper(c.PublicField)
	upper2 := strings.ToUpper(o.PublicField)

	return upper1 == upper2
}

// Basic -- go-cmp の基本パターンについてのサンプルです.
func Basic() error {
	// ---------------------------------------------------------------------------------
	// go-cmp は、比較処理用のライブラリ
	//
	// 標準の比較機能、つまり、reflect.DeepEqual よりもフレキシブルな比較処理を作ることが出来る
	//
	// 以下でインストールする。
	//   go get -u github.com/google/go-cmp/cmp
	//
	// よく利用する機能は、 cmp.Diff() と cmp.Equal()
	//
	// 基本的な機能として、対象となる型に Equal メソッドが定義されていれば
	// それを使って比較をしてくれる。
	//
	// ---
	// 参考
	//   - https://qiita.com/iszk/items/e799ec4d6f1f5eece706
	//   - https://pkg.go.dev/mod/github.com/google/go-cmp
	// ---------------------------------------------------------------------------------

	// ---------------------------------------------------------------------------------
	// cmp.Diff()
	//   diff コマンドのような比較結果を出力してくれる
	a1 := A{PublicField: "hello"}
	a2 := A{PublicField: "HeLLo"}

	if diff := cmp.Diff(a1, a2); diff != "" {
		// 差異あり
		output.Stdoutl("Diff(a1, a2)", diff)
	}

	// ---------------------------------------------------------------------------------
	// デフォルトでは プライベートなフィールド　があると panic する
	// そのまま実行すると以下のメッセージが出る
	//   panic: cannot handle unexported field at {gocmp.B}.privateField:
	//	  "github.com/devlights/try-golang/examples/advanced/gocmp".B
	b1 := B{
		A:            a1,
		privateField: "hello",
	}
	b2 := B{
		A:            a2,
		privateField: "world",
	}

	// ---------------------------------------------------------------------------------
	// プライベートなフィールドを比較対象に含めるには、オプションで設定する
	// cmp.AllowUnexported は内部で reflect.TypeOf して型情報を取得してマッピングを生成している
	opt := cmp.AllowUnexported(b1)
	if diff := cmp.Diff(b1, b2, opt); diff != "" {
		// 差異あり
		output.Stdoutl("Diff(b1, b2, cmp.AllowUnexported)", diff)
	}

	// プライベートなフィールドを無視するのも、オプションで設定する
	// cmpopts.IgnoreUnexported は内部で reflect.TypeOf して型情報を取得してマッピングを生成している
	opt = cmpopts.IgnoreUnexported(b1)
	if diff := cmp.Diff(b1, b2, opt); diff != "" {
		// 差異あり
		output.Stdoutl("Diff(b1, b2, cmpopts.IgnoreUnexported)", diff)
	}

	// ---------------------------------------------------------------------------------
	// Equal メソッドが定義されている場合、それを使ってくれる
	c1 := C{
		PublicField: "hello",
	}
	c2 := C{
		PublicField: "HeLLo",
	}

	if diff := cmp.Diff(c1, c2); diff != "" {
		// 差異あり
		output.Stdoutl("Diff(c1, c2)", diff)
	} else {
		output.Stdoutl("Diff(c1, c2)", "差異なし")
	}

	// ---------------------------------------------------------------------------------
	// cmp.Equal() は、Diff と同じように比較してくれるが 戻り値が bool となる
	output.Stdoutl("Equal(a1, a2)", cmp.Equal(a1, a2))
	output.Stdoutl("Equal(c1, c2)", cmp.Equal(c1, c2))

	// ---------------------------------------------------------------------------------
	// cmp.Comparer(f interface{}) Option を利用して、専用の Comparer を作って
	// 比較することも出来る
	//
	// cmp.Comparer() の引数は interface{} となっているが
	//   func (T,T) bool
	// を渡さないといけない。
	//  > The comparer f must be a function "func(T, T) bool"
	//  > and is implicitly filtered to input values assignable to T
	//
	// https://pkg.go.dev/github.com/google/go-cmp/cmp?tab=doc#Comparer
	//
	opt = cmp.Comparer(func(x, y A) bool {
		s1 := strings.ToLower(x.PublicField)
		s2 := strings.ToLower(y.PublicField)
		return s1 == s2
	})

	output.Stdoutl("Equal(a1, a2, cmp.Compare)", cmp.Equal(a1, a2, opt))

	return nil
}
