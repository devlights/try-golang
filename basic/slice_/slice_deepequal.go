package slice_

import (
	"fmt"
	"github.com/devlights/try-golang/lib/output"
	"reflect"
)

//SliceDeepEqual は、スライスに対して reflect.DeepEqual() した場合のサンプルです.
//
// REFERENCES::
//   - https://golang.org/ref/spec#Comparison_operators
//   - https://golang.org/pkg/reflect/#DeepEqual
func SliceDeepEqual() error {
	// ---------------------------------------------------------------------
	// スライスの比較について
	//
	// Goでは、そもそも slice, map, function を == で比較することができない
	// (https://golang.org/ref/spec#Comparison_operators)
	//
	// スライスはよく利用するデータタイプの一つであり、２つのシーケンスのデータが
	// 等しいかどうかをチェックすることもよくあることである。
	//
	// Goでは、その場合に reflect.DeepEqual() が利用できる.
	// (https://golang.org/pkg/reflect/#DeepEqual)
	// reflect.DeepEqual のドキュメントには詳細な説明が記載されている。
	// ---------------------------------------------------------------------
	sliceDataIsBaseType()
	sliceDataIsStruct()

	return nil
}

func sliceDataIsStruct() {

	output.Stdoutl("[Slice data is struct]", "-------------------------------------------")

	type (
		myData struct {
			name string
		}
	)

	// ---------------------------------------------------------------------
	// 同じ len, cap を持っている別々のスライスの場合
	// ---------------------------------------------------------------------
	var (
		sli1 = make([]myData, 0, 10)
		sli2 = make([]myData, 0, 10)
	)

	for i := 0; i < cap(sli1); i++ {
		s := fmt.Sprintf("%d", i)
		sli1 = append(sli1, myData{s})
		sli2 = append(sli2, myData{s})
	}

	output.Stdoutl("[sli1]", sli1)
	output.Stdoutl("[sli2]", sli2)
	output.Stdoutl("sli1 == sli2", reflect.DeepEqual(sli1, sli2))

	// ---------------------------------------------------------------------
	// 片方のスライスのデータ順序を変化させる
	// ---------------------------------------------------------------------
	swapper := reflect.Swapper(sli2)
	for i := 0; i < len(sli2)/2; i++ {
		swapper(i, (len(sli2)/2)+i)
	}

	output.Stdoutl("[sli1]", sli1)
	output.Stdoutl("[sli2]", sli2)
	output.Stdoutl("sli1 == sli2", reflect.DeepEqual(sli1, sli2))

	// ---------------------------------------------------------------------
	// 要素数が異なるスライスとの比較
	// ---------------------------------------------------------------------
	var (
		sli3 = make([]myData, 0, 5)
	)

	for i := 0; i < cap(sli3); i++ {
		s := fmt.Sprintf("%d", i)
		sli3 = append(sli3, myData{s})
	}

	output.Stdoutl("[sli1]", sli1)
	output.Stdoutl("[sli3]", sli3)
	output.Stdoutl("sli1 == sli3", reflect.DeepEqual(sli1, sli3))

	// ---------------------------------------------------------------------
	// sli1 の要素数を減らして, sli3 と同じにする
	// ---------------------------------------------------------------------
	sli1 = sli1[:cap(sli3)]

	output.Stdoutl("[sli1]", sli1)
	output.Stdoutl("[sli3]", sli3)
	output.Stdoutl("sli1 == sli3", reflect.DeepEqual(sli1, sli3))
}

func sliceDataIsBaseType() {

	output.Stdoutl("[Slice data is base type]", "-------------------------------------------")

	// ---------------------------------------------------------------------
	// 同じ len, cap を持っている別々のスライスの場合
	// ---------------------------------------------------------------------
	var (
		sli1 = make([]int, 0, 10)
		sli2 = make([]int, 0, 10)
	)

	for i := 0; i < cap(sli1); i++ {
		sli1 = append(sli1, i)
		sli2 = append(sli2, i)
	}

	output.Stdoutl("[sli1]", sli1)
	output.Stdoutl("[sli2]", sli2)
	output.Stdoutl("sli1 == sli2", reflect.DeepEqual(sli1, sli2))

	// ---------------------------------------------------------------------
	// 片方のスライスのデータ順序を変化させる
	// ---------------------------------------------------------------------
	swapper := reflect.Swapper(sli2)
	for i := 0; i < len(sli2)/2; i++ {
		swapper(i, (len(sli2)/2)+i)
	}

	output.Stdoutl("[sli1]", sli1)
	output.Stdoutl("[sli2]", sli2)
	output.Stdoutl("sli1 == sli2", reflect.DeepEqual(sli1, sli2))

	// ---------------------------------------------------------------------
	// 要素数が異なるスライスとの比較
	// ---------------------------------------------------------------------
	var (
		sli3 = make([]int, 0, 5)
	)

	for i := 0; i < cap(sli3); i++ {
		sli3 = append(sli3, i)
	}

	output.Stdoutl("[sli1]", sli1)
	output.Stdoutl("[sli3]", sli3)
	output.Stdoutl("sli1 == sli3", reflect.DeepEqual(sli1, sli3))

	// ---------------------------------------------------------------------
	// sli1 の要素数を減らして, sli3 と同じにする
	// ---------------------------------------------------------------------
	sli1 = sli1[:cap(sli3)]

	output.Stdoutl("[sli1]", sli1)
	output.Stdoutl("[sli3]", sli3)
	output.Stdoutl("sli1 == sli3", reflect.DeepEqual(sli1, sli3))
}
