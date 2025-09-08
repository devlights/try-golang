package sliceop

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
)

// Chunk は、slices.Chunk()のサンプルです。
func Chunk() error {
	//
	// Go1.23でサポートされたiterパッケージのサンプルとして
	// わざとマップからスライスへの変換をしている
	//
	var (
		months = map[string]string{
			"Janualy":   "1",
			"Febualy":   "2",
			"March":     "3",
			"April":     "4",
			"May":       "5",
			"June":      "6",
			"July":      "7",
			"August":    "8",
			"September": "9",
			"October":   "10",
			"November":  "11",
			"December":  "12",
		}
		monthNames  = slices.Collect(maps.Keys(months))
		maxLenMonth = slices.MaxFunc(monthNames, func(x, y string) int {
			return cmp.Compare(len(x), len(y))
		})
	)

	//
	// ソートして、３つずつ出力
	//
	const (
		nChunks = 3
	)

	slices.Sort(monthNames)
	for chunk := range slices.Chunk(monthNames, nChunks) {
		for v := range slices.Values(chunk) {
			fmt.Printf("%-*s", len(maxLenMonth), v)
		}

		fmt.Println("")
	}

	return nil

	/*
	   $ task
	   task: [build] go build -o "/home/dev/dev/github/try-golang/try-golang" .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: sliceop_iter_chunk

	   [Name] "sliceop_iter_chunk"
	   April    August   December
	   Febualy  Janualy  July
	   June     March    May
	   November October  September

	   [Elapsed] 81.717µs
	*/
}
