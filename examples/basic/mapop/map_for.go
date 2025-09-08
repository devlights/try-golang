package mapop

import "fmt"

// MapFor -- マップをループするサンプルです。
func MapFor() error {

	m := map[string]int{
		"apple":  100,
		"banana": 200,
	}

	// キーだけ欲しい場合
	for k := range m {
		fmt.Printf("KEY: %s\n", k)
	}

	// キーと値のペアが欲しい場合
	for k, v := range m {
		fmt.Printf("KEY: %s\tVALUE: %d\n", k, v)
	}

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: map_for

	   [Name] "map_for"
	   KEY: apple
	   KEY: banana
	   KEY: apple      VALUE: 100
	   KEY: banana     VALUE: 200


	   [Elapsed] 33.95µs
	*/

}
