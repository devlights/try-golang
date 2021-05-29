package runner

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/devlights/try-golang/mapping"
)

type (
	// Loop -- 実行をループ処理するコマンド
	Loop struct {
		Args *LoopArgs
	}

	// LoopArgs -- Loop の引数データを表します.
	LoopArgs struct {
		In      io.Reader              // 入力
		OneTime bool                   // 一回実行で完了するかどうか
		Mapping mapping.ExampleMapping // マッピング情報
	}
)

// NewLoopArgs -- 新しい LoopArgs を生成して返します.
func NewLoopArgs(in io.Reader, oneTime bool, m mapping.ExampleMapping) *LoopArgs {
	a := new(LoopArgs)
	a.In = in
	a.OneTime = oneTime
	a.Mapping = m
	return a
}

// NewLoop -- 新しい Loop を生成して返します.
func NewLoop(args *LoopArgs) *Loop {
	c := new(Loop)
	c.Args = args
	return c
}

// Run -- 実行します.
func (c *Loop) Run() error {
	var (
		in = c.Args.In
		oneTime = c.Args.OneTime
		mapping = c.Args.Mapping
	)

	fmt.Print("ENTER EXAMPLE NAME: ")

	stdinScanner := bufio.NewScanner(in)
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

		if oneTime {
			break
		}

	nextinput:
		fmt.Print("\n\n")
		fmt.Print("ENTER EXAMPLE NAME: ")
	}

	return nil
}

func (c *Loop) exec(target string, mapping mapping.ExampleMapping) error {
	execArgs := NewExecArgs(target, mapping)
	execCmd := NewExec(execArgs)

	if err := execCmd.Run(); err != nil {
		return err
	}

	return nil
}

func (c *Loop) makeCandidates(userInput string, mapping mapping.ExampleMapping) []string {
	candidates := make([]string, 0, len(mapping))
	for k := range mapping {
		key := string(k)
		if strings.Contains(key, userInput) {
			candidates = append(candidates, key)
		}
	}

	return candidates
}
