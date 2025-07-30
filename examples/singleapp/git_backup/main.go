package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type (
	AppArgs struct {
		NumLogCount int
		File        string
	}
)

var (
	appArgs AppArgs
	out     io.Writer
)

func init() {
	flag.IntVar(&appArgs.NumLogCount, "n", 1, "何世代前の版を取得するかの値")
	flag.StringVar(&appArgs.File, "f", "", "ファイル")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if appArgs.File == "" {
		log.Fatal("invalid argument: file")
	}

	out = os.Stdout

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// 1.パス文字が含まれていない場合、Prefixを取得してパス作る (git rev-parse --show-prefix)
	//   含んでいる場合、ユーザが明示的に指定しているとみなしPrefixは付けない
	// 2.SHA取得      (git log --pretty=format:"%h" -世代数 ファイル名)
	// 3.ファイル生成 (git show SHA:パス)
	//
	// MEMO:
	//   git log と git show では渡すパスの形が異なる
	//     - git log  は カレントディレクトリからの相対パスを受け付ける仕様
	//     - git show は リポジトリルートからのパスを受け付ける仕様

	var (
		git   = new(GitCmd)
		fpath = appArgs.File
		err   error
	)
	if !strings.Contains(fpath, "/") {
		var (
			prefix string
		)
		if prefix, err = git.Prefix(); err != nil {
			return err
		}

		fpath = filepath.Join(prefix, fpath)
	}

	var (
		sha string
	)
	if sha, err = git.Sha(appArgs.File, appArgs.NumLogCount); err != nil {
		return err
	}
	if sha == "" {
		return fmt.Errorf("SHA取得失敗: %s", appArgs.File)
	}

	var (
		r io.ReadCloser
	)
	if r, err = git.Show(sha, fpath); err != nil {
		return err
	}
	defer r.Close()

	if _, err = io.Copy(out, r); err != nil {
		return err
	}

	return nil
}
