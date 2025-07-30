package main

import (
	"io"
	"os/exec"
)

type (
	CmdReader struct {
		pipe io.ReadCloser
		cmd  *exec.Cmd
	}
)

var _ io.ReadCloser = (*CmdReader)(nil)

func (me *CmdReader) Read(p []byte) (int, error) {
	return me.pipe.Read(p)
}

func (me *CmdReader) Close() error {
	var (
		errPipe = me.pipe.Close()
		errWait = me.cmd.Wait()
	)
	if errPipe != nil {
		return errPipe
	}
	if errWait != nil {
		return errWait
	}

	return nil
}
