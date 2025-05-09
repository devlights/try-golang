package main

import (
	"log"

	"github.com/prometheus/procfs"
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		p   procfs.Proc
		err error
	)
	p, err = procfs.Self()
	if err != nil {
		return err
	}

	var (
		cmdline []string
	)
	cmdline, err = p.CmdLine()
	if err != nil {
		return err
	}

	log.Printf("[Self] pid=%d, cmdline=%v", p.PID, cmdline)

	return nil
}
