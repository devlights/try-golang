package main

import (
	"fmt"
	"github.com/devlights/try-golang/lib"
	"sort"
)

func printAllExampleNames(mapping lib.SampleMapping) {
	names := make([]string, 0, len(mapping))

	for k := range mapping {
		key := string(k)
		names = append(names, key)
	}

	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	fmt.Println("[Examples]")
	for _, v := range names {
		fmt.Printf("\t%s\n", v)
	}
}