package main

import (
	"bufio"
	"fmt"
	"github.com/devlights/try-golang/lib"
	"log"
	"os"
	"sort"
	"strings"
)

var (
	args    *Args
	mapping lib.SampleMapping
)

func init() {
	args = NewArgs()
	args.Parse()

	mapping = lib.NewSampleMapping()
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

func makeCandidates(userInput string) []string {
	candidates := make([]string, 0, len(mapping))
	for k := range mapping {
		if strings.Contains(k, userInput) {
			candidates = append(candidates, k)
		}
	}

	return candidates
}

func exec(target string) error {
	if v, ok := mapping[target]; ok {
		fmt.Printf("[Name] %q\n", target)
		if err := v(); err != nil {
			return err
		}
	}

	return nil
}

func runOnce(nameOfExample string) {
	if err := exec(nameOfExample); err != nil {
		log.Fatal(err)
	}
}

func runLoop() {
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

		candidates = makeCandidates(userInput)
		numberOfCandidate = len(candidates)

		switch {
		case numberOfCandidate == 0:
			fmt.Printf("Not found...Try Again")
			goto nextinput
		case numberOfCandidate == 1:
			userInput = candidates[0]
			if err := exec(userInput); err != nil {
				log.Fatal(err)
			}
		case 1 < numberOfCandidate:
			isPerfectMatchFound := false
			for _, c := range candidates {
				if c == userInput {
					runOnce(c)
					isPerfectMatchFound = true
					break
				}
			}

			if !isPerfectMatchFound {
				fmt.Printf("There's %d candidates found\n", len(candidates))

				sort.Slice(candidates, func(i, j int) bool {
					return candidates[i] < candidates[j]
				})

				for _, item := range candidates {
					fmt.Printf("\t%s\n", item)
				}

				goto nextinput
			}
		}

		if args.OneTime {
			break
		}

	nextinput:
		fmt.Print("\n\n")
		fmt.Print("ENTER EXAMPLE NAME: ")
	}
}

func main() {
	if args.ShowNames {
		printAllExampleNames()
		os.Exit(0)
	}

	defer fmt.Println("END")

	if args.ExampleName != "" {
		runOnce(args.ExampleName)
	} else {
		runLoop()
	}

	os.Exit(0)
}
