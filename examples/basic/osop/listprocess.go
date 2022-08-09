package osop

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/devlights/gomy/output"
)

// ListProcesses -- プロセスリストを取得するサンプルです.
//
// 本サンプルは Windows では動作しません。
//
// REFERENCES:
//   - https://stackoverflow.com/questions/9030680/list-of-currently-running-process-in-go
func ListProcesses() error {
	if runtime.GOOS == "windows" {
		return errors.New("sorry, this example doesn't run on windows")
	}

	// Linux 系のOSでは、 /proc ファイルシステムから動作しているプロセスの情報を取得することが出来る.
	// PIDは、 /proc/[pid]/exe が、実行ファイルへのシンボリックリンクとなっている.

	var (
		matches []string
		err     error
	)

	matches, err = filepath.Glob("/proc/*/exe")
	if err != nil {
		return err
	}

	for _, f := range matches {
		// シンボリックリンクの実体取得
		real, err := os.Readlink(f)
		if err != nil {
			// Permission denied は無視
			if errors.Is(err, os.ErrPermission) {
				continue
			}

			return err
		}

		pid := filepath.Base(filepath.Dir(f))

		output.Stdoutf("[info]", "name=%v\tpid=%v\n", real, pid)
	}

	return nil
}
