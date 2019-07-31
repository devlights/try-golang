package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/devlights/try-golang/lib"
	"log"
	"os"
	"sort"
	"strings"
)

var mapping = make(lib.SampleMapping)

func init() {
	mapping.MakeMapping()
}

func printAllExampleNames() {
	names := make([]string, 0, len(mapping))

	for k := range mapping {
		names = append(names, k)
	}

	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	fmt.Println("[Examples]")
	for _, v := range names {
		fmt.Printf("\t%s\n", v)
	}
}

func main() {
	var (
		onetime   = flag.Bool("onetime", false, "run only one time")
		showNames = flag.Bool("list", false, "show all example names")
	)

	flag.Parse()

	if *showNames {
		printAllExampleNames()
		return
	}

	fmt.Print("ENTER EXAMPLE NAME: ")

	stdinScanner := bufio.NewScanner(os.Stdin)
	for stdinScanner.Scan() {
		var (
			numberOfCandidate int
			candidates        []string
		)

		userInput := stdinScanner.Text()
		if userInput == "" {
			goto nextinput
		}

		if strings.ToLower(userInput) == "quit" {
			// 終了
			break
		}

		for k := range mapping {
			if strings.Contains(k, userInput) {
				candidates = append(candidates, k)
			}
		}

		numberOfCandidate = len(candidates)
		switch {
		case numberOfCandidate == 0:
			fmt.Printf("Not found...Try Again")
			goto nextinput
		case numberOfCandidate == 1:
			if v, ok := mapping[userInput]; ok {
				if err := v(); err != nil {
					log.Fatal(err)
				}
			}
		case 1 < numberOfCandidate:
			fmt.Printf("There's %d candidates found\n", len(candidates))

			for _, item := range candidates {
				fmt.Printf("\t%s\n", item)
			}

			goto nextinput
		}

		if *onetime {
			break
		}

	nextinput:
		fmt.Print("\n\n")
		fmt.Print("ENTER EXAMPLE NAME: ")
	}

	fmt.Println("END")
}
