package osop

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/devlights/gomy/output"
)

// Mkdir -- os.Mkdir/MkdirAllのサンプルです.
//
// REFERENCES:
//   - https://pkg.go.dev/os@go1.17.8
//   - https://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permissions
//   - https://stackoverflow.com/questions/33450980/how-to-remove-all-contents-of-a-directory-using-golang
func Mkdir() error {
	//
	// ディレクトリ作成時などにパーミッションを指定する必要があるが
	// ここには８進数のパーミッション値を直接指定することができる.
	//
	// ex: 0755
	//
	// でも、基本、os.ModePerm (0777) を指定しておけば良い。
	// そのシステムにおけるデフォルトのパーミッションは umask で制御されるべきであるため。
	//
	var (
		tmpDir  = os.TempDir()
		baseDir = filepath.Join(tmpDir, "try-golang", "osop")
		p1      = filepath.Join(baseDir, "d1")
		p2      = filepath.Join(baseDir, "d2")
	)

	err := os.RemoveAll(baseDir)
	if err != nil {
		return err
	}

	// 8進数指定
	err = os.MkdirAll(p1, 0755)
	if err != nil {
		return err
	}

	// os.ModePerm (0777) を指定
	err = os.MkdirAll(p2, os.ModePerm)
	if err != nil {
		return err
	}

	// 結果確認
	b, err := exec.Command("ls", "-l", baseDir).Output()
	if err != nil {
		return err
	}
	output.Stdoutf("[ls -l]", "%s\n", b)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: osop_mkdir

	   [Name] "osop_mkdir"
	   [ls -l]              total 0
	   drwxr-xr-x 2 gitpod gitpod 40 Jan 29 04:18 d1
	   drwxr-xr-x 2 gitpod gitpod 40 Jan 29 04:18 d2



	   [Elapsed] 2.049119ms
	*/

}
