package async

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// DirWalkRecursive は、非同期処理と再帰処理の組み合わせのサンプルです。
func DirWalkRecursive() error {
	var (
		wg = sync.WaitGroup{}  // 待ち合わせ用
		ch = make(chan string) // メイン処理と非同期処理との間でデータを受け渡すためのチャネル
	)

	// --------------------------------------------------
	// 再帰しながらディレクトリツリーを下り、情報を出力
	// --------------------------------------------------
	wg.Add(1)
	dir, _ := filepath.Abs(".")
	go listdir(dir, &wg, ch, 1)

	// --------------------------------------------------
	// 終わる時間は不定なため、再帰処理にデータを処理させながら
	// 同時に出力を実施。再帰処理完了とともに出力処理を止める。
	// --------------------------------------------------
	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	return nil
}

func listdir(dir string, wg *sync.WaitGroup, ch chan<- string, depth int) {
	defer wg.Done()

	var (
		chSubDirs = make([]chan string, 0)
		dprefix   = strings.Repeat("\t", depth-1) // ディレクトリ用のプレフィックス
		fprefix   = strings.Repeat("\t", depth)   // ファイル用のプレフィックス
	)

	// ディレクトリ名を出力
	d := dir
	if depth > 1 {
		d = filepath.Base(dir)
	}

	ch <- fmt.Sprintf("%s%s", dprefix, d)

	// --------------------------------------------------
	// 自身の配下を非同期で処理
	// --------------------------------------------------
	files, _ := os.ReadDir(dir)
	for _, f := range files {
		if f.IsDir() {
			// ドットで始まるディレクトリは無視
			if strings.HasPrefix(f.Name(), ".") {
				continue
			}

			// 配下に対して非同期で探索開始
			chSubDir := make(chan string)
			chSubDirs = append(chSubDirs, chSubDir)

			wgSubDir := sync.WaitGroup{}
			wgSubDir.Add(1)
			go listdir(filepath.Join(dir, f.Name()), &wgSubDir, chSubDir, depth+1)

			// 再帰処理を非同期で待ち合わせ
			go func() {
				wgSubDir.Wait()
				close(chSubDir)
			}()
		} else {
			// ファイルの場合は出力
			ch <- fmt.Sprintf("%s%s", fprefix, f.Name())
		}
	}

	// 配下の非同期処理を実施しながら、結果を親のチャネルに入れていく
	for _, subCh := range chSubDirs {
		for v := range subCh {
			ch <- v
		}
	}
}
