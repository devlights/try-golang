package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type (
	GitCmd struct{}
)

func (me *GitCmd) exec(args []string) (*CmdReader, error) {
	var (
		cmd    = exec.Command("git", args...)
		stdout io.ReadCloser
		err    error
	)
	if stdout, err = cmd.StdoutPipe(); err != nil {
		return nil, err
	}
	if err = cmd.Start(); err != nil {
		return nil, err
	}

	return &CmdReader{stdout, cmd}, nil
}

// Prefix は、git rev-parse --show-prefix を実行します。
func (me *GitCmd) Prefix() (string, error) {
	var (
		args   = []string{"rev-parse", "--show-prefix"}
		reader *CmdReader
		err    error
	)
	if reader, err = me.exec(args); err != nil {
		return "", err
	}
	defer reader.Close()

	var (
		buf = new(bytes.Buffer)
	)
	if _, err = io.Copy(buf, reader); err != nil {
		return "", err
	}

	return strings.ReplaceAll(buf.String(), "\n", ""), nil
}

// Sha は、 git log --pretty=format:"%h" を実行しcount番目のSHAを返します。
func (me *GitCmd) Sha(fpath string, count int) (string, error) {
	var (
		args   = []string{"log", "--pretty=format:'%h'", fpath}
		reader *CmdReader
		err    error
	)
	if reader, err = me.exec(args); err != nil {
		return "", err
	}
	defer reader.Close()

	var (
		buf = new(bytes.Buffer)
	)
	if _, err = io.Copy(buf, reader); err != nil {
		return "", err
	}

	var (
		scanner = bufio.NewScanner(buf)
		sha     string
	)
	for i := 0; scanner.Scan(); i++ {
		if i == count {
			sha = scanner.Text()
			break
		}
	}

	if err = scanner.Err(); err != nil {
		return "", err
	}

	return strings.ReplaceAll(sha, "'", ""), nil
}

// Show は、 git show sha:fpath を実行します。
func (me *GitCmd) Show(sha, fpath string) (io.ReadCloser, error) {
	var (
		args   = []string{"show", fmt.Sprintf("%s:%s", sha, fpath)}
		reader *CmdReader
		err    error
	)
	if reader, err = me.exec(args); err != nil {
		return nil, err
	}

	return reader, nil
}
