package cmdexec

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

// OsPipe は、(*Cmd).Stdout に os.Pipe の io.Writer を接続して処理するサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/os@go1.24.4#Pipe
func OsPipe() error {
	var (
		pr  *os.File
		pw  *os.File
		err error
	)
	pr, pw, err = os.Pipe()
	if err != nil {
		return err
	}
	defer pr.Close()

	var (
		name = "git"
		args = []string{"--no-pager", "log", "-m", "-r", "--name-only", "--pretty=raw", "-z"}
		cmd  = exec.Command(name, args...)
	)
	cmd.Stdout = pw
	if err = cmd.Start(); err != nil {
		pw.Close()
		return err
	}

	// 終了待機
	var (
		done = make(chan error, 1)
	)
	go func() {
		defer pw.Close()
		done <- cmd.Wait()
	}()

	const (
		MaxTokenSize = 1024 * 1024
	)
	var (
		scanner = bufio.NewScanner(pr)
		buf     = make([]byte, MaxTokenSize)
	)
	scanner.Buffer(buf, MaxTokenSize)

	for scanner.Scan() {
		log.Println(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	if err = <-done; err != nil {
		return err
	}

	return nil
}
