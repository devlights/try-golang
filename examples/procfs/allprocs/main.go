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
		fs  procfs.FS
		err error
	)
	fs, err = procfs.NewFS(procfs.DefaultMountPoint)
	if err != nil {
		return err
	}

	var (
		procs procfs.Procs // type Procs []Proc
	)
	procs, err = fs.AllProcs()
	if err != nil {
		return err
	}

	var (
		drop = func(v string, _ error) string { return v }
	)
	for i, p := range procs {
		log.Printf("[%02d] %7d: %s", i+1, p.PID, drop(p.Executable()))
	}

	return nil
}
