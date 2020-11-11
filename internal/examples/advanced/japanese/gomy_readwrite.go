package japanese

import (
	"io/ioutil"

	"github.com/devlights/gomy/fileio"
	"github.com/devlights/gomy/fileio/jp"
	"github.com/devlights/gomy/output"
)

// GomyReadWrite は、github.com/devlights/gomy の関数を利用するサンプルです。
func GomyReadWrite() error {
	dir, rmdirFn, err := fileio.TempDir("try-golang")
	if err != nil {
		return err
	}

	defer rmdirFn()

	fpath, err := func() (string, error) {
		name := fileio.TempFileName(dir, "try-golang")
		writer, releaseFn, ioErr := fileio.OpenWrite(name, jp.ShiftJis)
		if ioErr != nil {
			return "", nil
		}

		defer func() {
			_ = releaseFn()
		}()

		_, ioErr = writer.Write([]byte("こんにちわWorld"))
		if ioErr != nil {
			return "", nil
		}

		return name, nil
	}()

	if err != nil {
		return err
	}

	output.Stdoutl("fpath", fpath)

	err = func() error {
		reader, releaseFn, ioErr := fileio.OpenRead(fpath, jp.ShiftJis)
		if ioErr != nil {
			return ioErr
		}

		defer func() {
			_ = releaseFn()
		}()

		allData, ioErr := ioutil.ReadAll(reader)
		if ioErr != nil {
			return ioErr
		}

		output.Stdoutl("[sjis]", string(allData))

		return nil
	}()

	return nil
}
