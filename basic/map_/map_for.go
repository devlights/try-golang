package map_

import "fmt"

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
}
