package cmpop

import (
	"cmp"
	"slices"

	"github.com/devlights/gomy/output"
)

// CompositeSortKeys は、cmp.Or, cmp.Compareを用いて複合キーのソート処理を実装するサンプルです。
func CompositeSortKeys() error {
	type Person struct {
		Name string
		Age  uint8
	}

	var (
		people = []Person{
			{"Aikawa", 21},
			{"Tanaka", 22},
			{"Kato", 33},
			{"Suzuki", 44},
			{"Tanaka", 44},
			{"Aikawa", 66},
		}
	)

	output.Stdoutl("[before]", people)

	// 名前の昇順が第１キー、年齢の降順が第２キーとする
	slices.SortFunc(people, func(x, y Person) int {
		return cmp.Or(
			cmp.Compare(x.Name, y.Name),
			-cmp.Compare(x.Age, y.Age),
		)
	})

	output.Stdoutl("[after ]", people)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: cmpop_composite_sort_key

	   [Name] "cmpop_composite_sort_keys"
	   [before]             [{Aikawa 21} {Tanaka 22} {Kato 33} {Suzuki 44} {Tanaka 44} {Aikawa 66}]
	   [after ]             [{Aikawa 66} {Aikawa 21} {Kato 33} {Suzuki 44} {Tanaka 44} {Tanaka 22}]


	   [Elapsed] 62.25µs
	*/

}
