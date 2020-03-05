package struct_

import (
	"github.com/devlights/try-golang/lib/output"
	"reflect"
)

//StructDeepEqual は、構造体に対して reflect.DeepEqual() した場合のサンプルです.
//
// REFERENCES::
//   - https://golang.org/ref/spec#Comparison_operators
//   - https://golang.org/pkg/reflect/#DeepEqual
//   - https://qiita.com/Sekky0905/items/1ff4979d80b163e0aeb6
//   - https://stackoverflow.com/questions/24534072/how-to-compare-if-two-structs-slices-or-maps-are-equal
//   - https://github.com/google/go-cmp
func StructDeepEqual() error {
	// ---------------------------------------------------------------------
	// 構造体の比較について
	//
	// Goでは、 == は等値ではなく等価となる。
	// しかし、構造体の内部で持っているポインタの先までは見に行ってくれるわけではない。
	//
	// Goでは、その場合に reflect.DeepEqual() が利用できる.
	// (https://golang.org/pkg/reflect/#DeepEqual)
	// reflect.DeepEqual のドキュメントには詳細な説明が記載されている。
	//
	// なお、reflect.DeepEqual() にも少し難点があって、この関数は構造体の非公開フィールドも
	// 比較対象として処理してしまう仕様となっている。有名なのが Time型 で同じ時間を示しているのに
	// reflect.DeepEqual() で比較すると false になってしまう場合がある。
	// (https://stackoverflow.com/questions/24534072/how-to-compare-if-two-structs-slices-or-maps-are-equal)
	//
	// これが困る場合は、 go-cmp (https://github.com/google/go-cmp) のようなライブラリを利用することでうまくいく場合がある
	// go-cmp は、デフォルトで非公開フィールドを見ず、また Equal という名前のメソッドを
	// 実装していれば、それを呼び出してくれる機能も持っている。
	// ---------------------------------------------------------------------
	type (
		// 別の構造体に組み込んで利用されるための構造体
		baseSt struct {
			BaseField           int
			unExportedBaseField int
		}

		refSt struct {
			RefField int
		}

		st struct {
			baseSt
			Field    int
			RefField *refSt
		}
	)

	var (
		st1 = st{
			baseSt:   baseSt{BaseField: 1, unExportedBaseField: 1},
			Field:    1,
			RefField: &refSt{RefField: 1},
		}
		st2 = st{
			baseSt:   baseSt{BaseField: 1, unExportedBaseField: 1},
			Field:    1,
			RefField: &refSt{RefField: 1},
		}
	)

	// ---------------------------------------------------------------------
	// ２つのデータが等値ではないことを確認
	// ---------------------------------------------------------------------
	output.Stdoutl("===> 等値かどうかのチェック")
	output.Stdoutf("[&st1]", "%p\n", &st1)
	output.Stdoutf("[&st2]", "%p\n", &st2)
	output.Stdoutl("[&st1 == &st2]", &st1 == &st2)

	// ---------------------------------------------------------------------
	// フィールドの値は全く同じになっている
	// 唯一、RefFieldに設定されているポインタのアドレスが異なる状態
	//
	// == での比較は、ポインタの先までは見に行かないので false となるが
	// reflect.DeepEqual() は、ポインタの先まで見に行って比較するので true となる
	// ---------------------------------------------------------------------
	output.Stdoutl("===> 等価かどうかのチェック (1)")
	output.Stdoutl("   ===> 全てのフィールドの値が同じ")
	output.Stdoutl("[st1]", st1)
	output.Stdoutl("[st2]", st2)
	output.Stdoutl("[st1 == st2]", st1 == st2)
	output.Stdoutl("[deepequal(st1,st2)]", reflect.DeepEqual(st1, st2))

	// ---------------------------------------------------------------------
	// 片方の構造体のデータを変化させる (自身のフィールドのみを変化)
	// ---------------------------------------------------------------------
	st2.Field = 2

	output.Stdoutl("===> 等価かどうかのチェック (2)")
	output.Stdoutl("   ===> 片方の構造体のデータを変化させる (自身のフィールドのみを変化)")
	output.Stdoutl("[st1]", st1)
	output.Stdoutl("[st2]", st2)
	output.Stdoutl("[st1 == st2]", st1 == st2)
	output.Stdoutl("[deepequal(st1,st2)]", reflect.DeepEqual(st1, st2))

	// ---------------------------------------------------------------------
	// 片方の構造体のデータを変化させる (組み込み構造体側の公開フィールド値を変化)
	// ---------------------------------------------------------------------
	st2.Field = 1
	st2.BaseField = 2

	output.Stdoutl("===> 等価かどうかのチェック (3)")
	output.Stdoutl("   ===> 片方の構造体のデータを変化させる (組み込み構造体側の公開フィールド値を変化)")
	output.Stdoutl("[st1]", st1)
	output.Stdoutl("[st2]", st2)
	output.Stdoutl("[st1 == st2]", st1 == st2)
	output.Stdoutl("[deepequal(st1,st2)]", reflect.DeepEqual(st1, st2))

	// ---------------------------------------------------------------------
	// 片方の構造体のデータを変化させる (組み込み構造体側の非公開フィールド値を変化)
	// ---------------------------------------------------------------------
	st2.BaseField = 1
	st2.unExportedBaseField = 2

	output.Stdoutl("===> 等価かどうかのチェック (4)")
	output.Stdoutl("   ===> 片方の構造体のデータを変化させる (組み込み構造体側の非公開フィールド値を変化)")
	output.Stdoutl("[st1]", st1)
	output.Stdoutl("[st2]", st2)
	output.Stdoutl("[st1 == st2]", st1 == st2)
	output.Stdoutl("[deepequal(st1,st2)]", reflect.DeepEqual(st1, st2))

	// ---------------------------------------------------------------------
	// 片方の構造体がポインタで保持している先のデータのフィールドを変化させる
	// ---------------------------------------------------------------------
	st2.unExportedBaseField = 1
	st2.RefField.RefField = 2

	output.Stdoutl("===> 等価かどうかのチェック (5)")
	output.Stdoutl("   ===> 片方の構造体がポインタで保持している先のデータのフィールドを変化させる")
	output.Stdoutl("[st1]", st1)
	output.Stdoutl("[st2]", st2)
	output.Stdoutl("[st1 == st2]", st1 == st2)
	output.Stdoutl("[deepequal(st1,st2)]", reflect.DeepEqual(st1, st2))

	// ---------------------------------------------------------------------
	// 片方の構造体がポインタで保持しているデータ自体を変更する。
	// しかし、データのフィールド値は同じ.
	//
	// == での比較は、ポインタの先までは見に行かないので false となるが
	// reflect.DeepEqual() は、ポインタの先まで見に行って比較するので true となる
	// ---------------------------------------------------------------------
	st2.RefField = &refSt{RefField: 1}

	output.Stdoutl("===> 等価かどうかのチェック (6)")
	output.Stdoutl("   ===> 片方の構造体がポインタで保持しているデータ自体を変更")
	output.Stdoutl("   ===> しかし、データのフィールド値は同じ")
	output.Stdoutl("[st1]", st1)
	output.Stdoutl("[st2]", st2)
	output.Stdoutl("[st1 == st2]", st1 == st2)
	output.Stdoutl("[deepequal(st1,st2)]", reflect.DeepEqual(st1, st2))

	return nil
}
