package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var (
	tmpDir string
)

func TestMain(m *testing.M) {
	ret := m.Run()

	// テスト終了後に一時ディレクトリが削除されているか確認
	if tmpDir != "" {
		_, err := os.Stat(tmpDir)
		fmt.Printf("[teardown] os.Stat() ==> %v", err)
	}

	os.Exit(ret)
}

func TestTempDir(t *testing.T) {
	//
	// t.TempDir() にてテスト時に利用できる一時ディレクトリの
	// パスが取得出来る。この一時ディレクトリはテスト時に作成されて
	// テスト終了後に自動的に消去される。テスト時に手動で一時ディレクトリを
	// 確保しておく手間が無くなるため、とても便利。
	//
	tmpDir = t.TempDir()
	t.Logf("tmpDir=%s", tmpDir)

	fi, err := os.Stat(tmpDir)
	t.Logf("IsDir=%v, Name=%s, Err=%v", fi.IsDir(), fi.Name(), err)

	// ディレクトリが存在しているだけでは何なので
	// 何かのファイルを書き込んでおく
	p := filepath.Join(tmpDir, "hello.txt")
	os.WriteFile(p, []byte("hello world"), 0777)

	fi, err = os.Stat(p)
	t.Logf("IsDir=%v, Name=%s, Err=%v", fi.IsDir(), fi.Name(), err)

	// t.TempDir() の呼び出しは何回でも良い
	// その度に、異なるディレクトリが返る
	for i := 0; i < 10; i++ {
		t.Logf("tmpDir=%s", t.TempDir())
	}
}
