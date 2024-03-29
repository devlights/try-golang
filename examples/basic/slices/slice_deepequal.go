package slices

import (
	"fmt"
	"reflect"

	"github.com/devlights/gomy/output"
)

// DeepEqual は、スライスに対して reflect.DeepEqual() した場合のサンプルです.
//
// REFERENCES::
//   - https://golang.org/ref/spec#Comparison_operators
//   - https://golang.org/reflect/#DeepEqual
//   - https://qiita.com/Sekky0905/items/1ff4979d80b163e0aeb6
//   - https://medium.com/golangspec/comparison-operators-in-go-910d9d788ec0
func DeepEqual() error {
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
	// (https://golang.org/reflect/#DeepEqual)
	// reflect.DeepEqual のドキュメントには詳細な説明が記載されている。
	//
	// Goの場合、等値を確認する場合はポインタを比較する
	//   &sli1 == &sli2
	// 等価を確認する場合に、reflect.DeepEqual() を利用する
	// ---------------------------------------------------------------------
	sliceDataIsBaseType()
	sliceDataIsStruct()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: slice_deep_equal

	   [Name] "slice_deep_equal"
	   [Slice data is base type] -------------------------------------------
	   [&sli1]              0xc0000125e8
	   [&sli2]              0xc000012600
	   &sli1 == &sli2       false
	   [sli1]               [0 1 2 3 4 5 6 7 8 9]
	   [sli2]               [0 1 2 3 4 5 6 7 8 9]
	   sli1 == sli2         true
	   [sli1]               [0 1 2 3 4 5 6 7 8 9]
	   [sli2]               [5 6 7 8 9 0 1 2 3 4]
	   sli1 == sli2         false
	   [&sli3]              0xc0000126f0
	   &sli1 == &sli3       false
	   [sli1]               [0 1 2 3 4 5 6 7 8 9]
	   [sli3]               [0 1 2 3 4]
	   sli1 == sli3         false
	   [sli1]               [0 1 2 3 4]
	   [sli3]               [0 1 2 3 4]
	   sli1 == sli3         true
	   [Slice data is struct] -------------------------------------------
	   [&sli1]              0xc0000127c8
	   [&sli2]              0xc0000127e0
	   &sli1 == &sli2       false
	   [sli1]               [{0} {1} {2} {3} {4} {5} {6} {7} {8} {9}]
	   [sli2]               [{0} {1} {2} {3} {4} {5} {6} {7} {8} {9}]
	   sli1 == sli2         true
	   [sli1]               [{0} {1} {2} {3} {4} {5} {6} {7} {8} {9}]
	   [sli2]               [{5} {6} {7} {8} {9} {0} {1} {2} {3} {4}]
	   sli1 == sli2         false
	   [&sli3]              0xc0000128d0
	   &sli1 == &sli3       false
	   [sli1]               [{0} {1} {2} {3} {4} {5} {6} {7} {8} {9}]
	   [sli3]               [{0} {1} {2} {3} {4}]
	   sli1 == sli3         false
	   [sli1]               [{0} {1} {2} {3} {4}]
	   [sli3]               [{0} {1} {2} {3} {4}]
	   sli1 == sli3         true


	   [Elapsed] 381.93µs
	*/

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

	output.Stdoutf("[&sli1]", "%p\n", &sli1)
	output.Stdoutf("[&sli2]", "%p\n", &sli2)
	output.Stdoutl("&sli1 == &sli2", &sli1 == &sli2)

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

	output.Stdoutf("[&sli3]", "%p\n", &sli3)
	output.Stdoutl("&sli1 == &sli3", &sli1 == &sli3)

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

	output.Stdoutf("[&sli1]", "%p\n", &sli1)
	output.Stdoutf("[&sli2]", "%p\n", &sli2)
	output.Stdoutl("&sli1 == &sli2", &sli1 == &sli2)

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

	output.Stdoutf("[&sli3]", "%p\n", &sli3)
	output.Stdoutl("&sli1 == &sli3", &sli1 == &sli3)

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
