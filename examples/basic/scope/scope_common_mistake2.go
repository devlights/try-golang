package scope

import (
	"os"

	"github.com/devlights/gomy/output"
)

var (
	_cwd2 string
)

func loadcwd2() error {
	// ローカルスコープで _cwd2 が生成されるのを防ぐために
	// err を var宣言 で先に宣言し、:= を使わないようにする
	var err error

	_cwd2, err = os.Getwd()
	if err != nil {
		return err
	}

	output.Stdoutl("[loadcwd]", _cwd2)

	return nil
}

// CommonMistake2 -- CommonMistake1の間違い修正パターン (1)
func CommonMistake2() error {
	if err := loadcwd2(); err != nil {
		return err
	}

	output.Stdoutl("[main]", _cwd2)

	return nil
}
