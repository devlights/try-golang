package main

import (
	"flag"
	"log"
	"os"
	"runtime/debug"
)

type subcommand string

const (
	subcmd1 subcommand = "cmd1"
	subcmd2 subcommand = "cmd2"
	subver  subcommand = "version"
	subhelp subcommand = "help"
)

// cmd1 subcommand
var (
	cmd1     = flag.NewFlagSet("cmd1", flag.ExitOnError)
	cmd1OptC = cmd1.Int("c", 0, "option c")
)

// cmd2 subcommand
var (
	cmd2     = flag.NewFlagSet("cmd2", flag.ExitOnError)
	cmd2OptF = cmd2.String("f", "", "option f")
)

func init() {
	log.SetFlags(0)
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	var (
		subcmd = subcommand(os.Args[1])
		args   = os.Args[2:]
	)

	switch subcmd {
	case subcmd1:
		cmd1.Parse(args)
		runCmd1()
	case subcmd2:
		cmd2.Parse(args)
		runCmd2()
	case subver:
		version()
		return
	case subhelp:
		help()
		return
	default:
		help()
		os.Exit(1)
	}
}

func runCmd1() {
	log.Printf("option c is %v", *cmd1OptC)
}

func runCmd2() {
	log.Printf("option f is %v", *cmd2OptF)
}

func version() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Println("error: call debug.ReadBuildInfo()")
		os.Exit(1)
	}

	for _, s := range info.Settings {
		if s.Key == "vcs.revision" {
			log.Printf("version: vX.Y.Z (%s)", s.Value)
			return
		}
	}
}

func help() {
	log.Println("Usage: app <cmd1|cmd2|version|help>")
	cmd1.Usage()
	cmd2.Usage()
}
