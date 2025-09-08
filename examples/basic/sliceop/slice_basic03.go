package sliceop

import "fmt"

// Basic03 -- スライスについてのサンプル
func Basic03() error {
	// REFERENCES:: http://bit.ly/2W6PsVM
	// GO言語には、python の del l[1] のように
	// サクッと要素を削除する方法が存在しない

	// 速度は速いが、リストの並びが一部変わってしまうパターン
	listA := []string{"A", "B", "C", "D", "E"}
	index := 2

	listA[index] = listA[len(listA)-1]
	listA[len(listA)-1] = ""
	listA = listA[:len(listA)-1]

	fmt.Println("ListA: ", listA)

	// 速度は遅いが、並び順を保つパターン
	listB := []string{"A", "B", "C", "D", "E"}

	copy(listB[index:], listB[index+1:])
	listB[len(listB)-1] = ""
	listB = listB[:len(listB)-1]

	fmt.Println("ListB: ", listB)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: slice_basic03

	   [Name] "slice_basic03"
	   ListA:  [A B E D]
	   ListB:  [A B D E]


	   [Elapsed] 36.15µs
	*/

}
