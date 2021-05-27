package scope

import (
	"os"

	"github.com/devlights/gomy/output"
)

type (
	_pkginfo struct {
		cwd string
	}
)

var (
	pkginfo _pkginfo
)

func loadcwd3() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	pkginfo.cwd = cwd
	output.Stdoutl("[loadcwd]", pkginfo.cwd)

	return nil
}

// CommonMistake3 -- CommonMistake1の間違い修正パターン (2)
func CommonMistake3() error {
	if err := loadcwd3(); err != nil {
		return err
	}

	output.Stdoutl("[main]", pkginfo.cwd)

	return nil
}
