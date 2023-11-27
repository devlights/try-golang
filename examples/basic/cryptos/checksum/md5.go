package checksum

import (
	"crypto/md5"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/devlights/gomy/output"
)

// Md5Checksum -- crypto/md5 のサンプルです.
func Md5Checksum() error {
	const (
		goExtension = ".go"
	)

	var (
		excludeTargets = []func(p string, i os.FileInfo) bool{
			func(p string, i os.FileInfo) bool { return i.IsDir() },
			func(p string, i os.FileInfo) bool { return i.Size() == 0 },
			func(p string, i os.FileInfo) bool { return !strings.HasSuffix(p, goExtension) },
		}
	)

	wd, _ := os.Getwd()
	output.Stdoutl("[working directory]", wd)

	err := filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, isExcludeTarget := range excludeTargets {
			if isExcludeTarget(path, info) {
				return nil
			}
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		// md5.New() は hash.Hash を返す
		// hash.Hash は、io.Writer を内包しているので io.Copy でデータを書き込む
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			return nil
		}

		// チェックサムのサイズは [md5.Size]byte となっている
		checksum := h.Sum(nil)[:md5.Size]
		output.Stdoutf("[checksum]", "%s\t%x\n", path, checksum)

		return nil
	})

	return err

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: crypto_md5_checksum

	   [Name] "crypto_md5_checksum"
	   [working directory]  /workspace/try-golang
	   [checksum]           /workspace/try-golang/builder/builder.go   988b24083283afcc9d88267266dc4400
	   [checksum]           /workspace/try-golang/builder/builder_test.go      2d96e2295bd02557a23b6fc7510a43e6
	   [checksum]           /workspace/try-golang/builder/doc.go       1a37ae8953bfae6ebeb675b990c53811
	   [checksum]           /workspace/try-golang/cmd/args.go  f0e4e8fcc1be9f1491d464c5276f0573
	   [checksum]           /workspace/try-golang/cmd/root.go  8e975214335c7e09f8ba312de7d64f29

	   snip...

	   [Elapsed] 40.861369ms
	*/

}
