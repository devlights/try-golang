package sync_and_close

import (
	"os"

	"github.com/devlights/gomy/output"
)

// SyncAndClose は、ファイルを扱う際にクローズする前にSyncしてからCloseするサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/os@go1.20.4#File.Sync
//   - https://pkg.go.dev/os@go1.20.4#File.Close
//   - https://yasukata.hatenablog.com/entry/2020/06/24/072609
//   - https://yasukata.hatenablog.com/entry/2020/06/23/031622
func SyncAndClose() error {
	//
	// ファイルに書き込む際、通常はCloseの呼び出しをdeferで書いて
	// エラー処理を省略してしまうことが多いが、ちゃんとファイルに書き出せたか
	// どうかを判定する場合は Sync() を呼び出す方が良い。
	// Sync() は、内部でシステムコール fsync の呼び出しを行う。
	//

	//
	// Write
	//
	var (
		fpath string
		err   error
	)

	fpath, err = write()
	if err != nil {
		return err
	}
	defer os.Remove(fpath)

	//
	// Read
	//
	var (
		data []byte
	)

	data, err = read(fpath)
	if err != nil {
		return err
	}

	output.Stdoutl("[Data]", string(data))

	return nil
}

func write() (string, error) {
	// Create
	var (
		file *os.File
		err  error
	)

	file, err = os.CreateTemp(os.TempDir(), "trygolang-")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Write
	var (
		buf = []byte("hello world")
	)

	_, err = file.Write(buf)
	if err != nil {
		return "", err
	}

	// Sync (Flush)
	//   - https://yasukata.hatenablog.com/entry/2020/06/24/072609
	err = file.Sync()
	if err != nil {
		return file.Name(), err
	}

	return file.Name(), nil
}

func read(fpath string) ([]byte, error) {
	// Read
	var (
		data []byte
		err  error
	)

	data, err = os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	return data, nil
}
