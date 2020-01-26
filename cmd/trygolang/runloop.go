package main

import (
	"bufio"
	"fmt"
	"github.com/devlights/try-golang/lib"
	"os"
	"sort"
	"strings"
)

type (
	RunLoopCommand struct {
		Args *RunLoopArgs
	}

	RunLoopArgs struct {
		MainArgs *Args
		Mapping  lib.SampleMapping
	}
)

func NewRunLoopArgs(mainArgs *Args, mapping lib.SampleMapping) *RunLoopArgs {
	a := new(RunLoopArgs)
	a.MainArgs = mainArgs
	a.Mapping = mapping
	return a
}

func NewRunLoopCommand(args *RunLoopArgs) *RunLoopCommand {
	c := new(RunLoopCommand)
	c.Args = args
	return c
}

func (c *RunLoopCommand) Run() error {
	var (
		mainArgs = c.Args.MainArgs
		mapping  = c.Args.Mapping
	)

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
			break
		}

		candidates = c.makeCandidates(userInput, mapping)
		numberOfCandidate = len(candidates)

		switch {
		case numberOfCandidate == 0:
			fmt.Printf("Not found...Try Again")
			goto nextinput
		case numberOfCandidate == 1:
			userInput = candidates[0]

			if err := c.exec(userInput, mapping); err != nil {
				return err
			}
		case 1 < numberOfCandidate:
			isPerfectMatchFound := false
			for _, candidate := range candidates {
				if candidate == userInput {
					if err := c.exec(candidate, mapping); err != nil {
						return err
					}

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

		if mainArgs.OneTime {
			break
		}

	nextinput:
		fmt.Print("\n\n")
		fmt.Print("ENTER EXAMPLE NAME: ")
	}

	return nil
}

func (c *RunLoopCommand) exec(target string, mapping lib.SampleMapping) error {
	execArgs := NewExecArgs(target, mapping)
	execCmd := NewExecCommand(execArgs)

	if err := execCmd.Run(); err != nil {
		return err
	}

	return nil
}

func (c *RunLoopCommand) makeCandidates(userInput string, mapping lib.SampleMapping) []string {
	candidates := make([]string, 0, len(mapping))
	for k := range mapping {
		if strings.Contains(k, userInput) {
			candidates = append(candidates, k)
		}
	}

	return candidates
}
