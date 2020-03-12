package maps

import (
	"reflect"
	"strconv"

	"github.com/devlights/gomy/output"
)

// MapDeepEqual は、マップに対して reflect.DeepEqual() した場合のサンプルです.
//
// REFERENCES::
//   - https://golang.org/ref/spec#Comparison_operators
//   - https://golang.org/pkg/reflect/#DeepEqual
//   - https://qiita.com/Sekky0905/items/1ff4979d80b163e0aeb6
//   - https://medium.com/golangspec/comparison-operators-in-go-910d9d788ec0
func MapDeepEqual() error {
	// ---------------------------------------------------------------------
	// マップの比較について
	//
	// Goでは、そもそも slice, map, function を == で比較することができない
	// (https://golang.org/ref/spec#Comparison_operators)
	//
	// とはいえ、マップ同士を比較したいときもたまにはある。
	//
	// Goでは、その場合に reflect.DeepEqual() が利用できる.
	// (https://golang.org/pkg/reflect/#DeepEqual)
	// reflect.DeepEqual のドキュメントには詳細な説明が記載されている。
	//
	// Goの場合、等値を確認する場合はポインタを比較する
	//   &m1 == &m2
	// 等価を確認する場合に、reflect.DeepEqual() を利用する
	// ---------------------------------------------------------------------
	mapDataIsBaseType()
	mapDataIsStruct()

	return nil
}

func mapDataIsBaseType() {

	output.Stdoutl("[Map data is base type]", "-------------------------------------------")

	// ---------------------------------------------------------------------
	// 同じ Key, Value を持っている別々のマップの場合
	// ---------------------------------------------------------------------
	var (
		m1 = make(map[int]string)
		m2 = make(map[int]string)
	)

	output.Stdoutf("[&m1]", "%p\n", &m1)
	output.Stdoutf("[&m2]", "%p\n", &m2)
	output.Stdoutl("&m1 == &m2", &m1 == &m2)

	for i := 0; i < 10; i++ {
		m1[i] = strconv.Itoa(i)
		m2[i] = strconv.Itoa(i)
	}

	output.Stdoutl("[m1]", m1)
	output.Stdoutl("[m2]", m2)
	output.Stdoutl("m1 == m2", reflect.DeepEqual(m1, m2))

	// ---------------------------------------------------------------------
	// 片方のマップのデータを変化させる
	// ---------------------------------------------------------------------
	for i := 0; i < 10; i++ {
		m2[i] = strconv.Itoa(i + (100 - i))
	}

	output.Stdoutl("[m1]", m1)
	output.Stdoutl("[m2]", m2)
	output.Stdoutl("m1 == m2", reflect.DeepEqual(m1, m2))

	// ---------------------------------------------------------------------
	// 要素数が異なるマップと比較
	// ---------------------------------------------------------------------
	var (
		m3 = make(map[int]string)
	)

	output.Stdoutf("[&m3]", "%p\n", &m3)
	output.Stdoutl("&m1 == &m3", &m1 == &m3)

	for i := 0; i < 5; i++ {
		m3[i] = strconv.Itoa(i)
	}

	output.Stdoutl("[m1]", m1)
	output.Stdoutl("[m3]", m3)
	output.Stdoutl("m1 == m3", reflect.DeepEqual(m1, m3))

	// ---------------------------------------------------------------------
	// m1の要素数を減らして、m3と同じにして比較
	// ---------------------------------------------------------------------
	for i := 5; i < 10; i++ {
		delete(m1, i)
	}

	output.Stdoutl("[m1]", m1)
	output.Stdoutl("[m3]", m3)
	output.Stdoutl("m1 == m3", reflect.DeepEqual(m1, m3))
}

func mapDataIsStruct() {

	output.Stdoutl("[Map data is struct]", "-------------------------------------------")

	type (
		myKey struct {
			id int
		}

		myData struct {
			name string
		}
	)

	// ---------------------------------------------------------------------
	// 同じ Key, Value を持っている別々のマップの場合
	// ---------------------------------------------------------------------
	var (
		m1 = make(map[myKey]myData)
		m2 = make(map[myKey]myData)
	)

	output.Stdoutf("[&m1]", "%p\n", &m1)
	output.Stdoutf("[&m2]", "%p\n", &m2)
	output.Stdoutl("&m1 == &m2", &m1 == &m2)

	for i := 0; i < 10; i++ {
		m1[myKey{i}] = myData{strconv.Itoa(i)}
		m2[myKey{i}] = myData{strconv.Itoa(i)}
	}

	output.Stdoutl("[m1]", m1)
	output.Stdoutl("[m2]", m2)
	output.Stdoutl("m1 == m2", reflect.DeepEqual(m1, m2))

	// ---------------------------------------------------------------------
	// 片方のマップのデータを変化させる
	// ---------------------------------------------------------------------
	for i := 0; i < 10; i++ {
		m2[myKey{i}] = myData{strconv.Itoa(i + (100 - i))}
	}

	output.Stdoutl("[m1]", m1)
	output.Stdoutl("[m2]", m2)
	output.Stdoutl("m1 == m2", reflect.DeepEqual(m1, m2))

	// ---------------------------------------------------------------------
	// 要素数が異なるマップと比較
	// ---------------------------------------------------------------------
	var (
		m3 = make(map[myKey]myData)
	)

	output.Stdoutf("[&m3]", "%p\n", &m3)
	output.Stdoutl("&m1 == &m3", &m1 == &m3)

	for i := 0; i < 5; i++ {
		m3[myKey{i}] = myData{strconv.Itoa(i)}
	}

	output.Stdoutl("[m1]", m1)
	output.Stdoutl("[m3]", m3)
	output.Stdoutl("m1 == m3", reflect.DeepEqual(m1, m3))

	// ---------------------------------------------------------------------
	// m1の要素数を減らして、m3と同じにして比較
	// ---------------------------------------------------------------------
	for i := 5; i < 10; i++ {
		delete(m1, myKey{i})
	}

	output.Stdoutl("[m1]", m1)
	output.Stdoutl("[m3]", m3)
	output.Stdoutl("m1 == m3", reflect.DeepEqual(m1, m3))
}
