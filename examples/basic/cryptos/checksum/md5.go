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
}
