package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	DefaultCommitFrom      = "HEAD~1"
	DefaultCommitTo        = "HEAD"
	DefaultRepoPath        = "."
	DefaultArchivePrefix   = "archive"
	DefaultArchiveFilePath = "./archive.zip"
)

type (
	Gitcmd string
)

func NewGitcmd() Gitcmd {
	return Gitcmd("git")
}

func (me Gitcmd) Exec(args ...string) ([]byte, error) {
	var (
		cmd = exec.Command(string(me), args...)
		out []byte
		err error
	)
	out, err = cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return bytes.TrimSuffix(out, []byte("\n")), nil
}

func (me Gitcmd) ExecString(args ...string) (string, error) {
	var (
		out []byte
		err error
	)
	out, err = me.Exec(args...)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

type (
	Args struct {
		CommitFrom      string
		CommitTo        string
		RepoPath        string
		ArchivePrefix   string
		ArchiveFilePath string
		Verbose         bool
	}
)

func (me *Args) restore() {
	if me.CommitFrom == "" {
		me.CommitFrom = DefaultCommitFrom
	}
	if me.CommitTo == "" {
		me.CommitTo = DefaultCommitTo
	}
	if me.RepoPath == "" {
		me.RepoPath = DefaultRepoPath
	}
	if me.ArchivePrefix == "" {
		me.ArchivePrefix = DefaultArchivePrefix
	}
	if me.ArchiveFilePath == "" {
		me.ArchiveFilePath = DefaultArchiveFilePath
	}
}

var (
	args Args
)

func init() {
	flag.StringVar(&args.CommitFrom, "from", DefaultCommitFrom, "起点となるコミットハッシュ")
	flag.StringVar(&args.CommitTo, "to", DefaultCommitTo, "終点となるコミットハッシュ")
	flag.StringVar(&args.RepoPath, "repo", DefaultRepoPath, "リポジトリパス")
	flag.StringVar(&args.ArchivePrefix, "prefix", DefaultArchivePrefix, "アーカイブ内のプレフィックスディレクトリ")
	flag.StringVar(&args.ArchiveFilePath, "archive", DefaultArchiveFilePath, "アーカイブファイル")
	flag.BoolVar(&args.Verbose, "v", false, "詳細表示")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	(&args).restore()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		git = NewGitcmd()
		err error
	)

	// リポジトリの場所に移動
	err = os.Chdir(args.RepoPath)
	if err != nil {
		return err
	}

	// 最新のコミットハッシュを取得
	var (
		from string
		to   string
	)
	from, err = git.ExecString("rev-parse", args.CommitFrom)
	if err != nil {
		return err
	}
	to, err = git.ExecString("rev-parse", args.CommitTo)
	if err != nil {
		return err
	}

	if args.Verbose {
		log.Printf("起点コミットハッシュ： %s(%s)", args.CommitFrom, from)
		log.Printf("終点コミットハッシュ： %s(%s)", args.CommitTo, to)
	}

	// 変更されたファイルの一覧を取得
	var (
		diff  string
		files []string
	)
	diff, err = git.ExecString("diff", "--name-only", from, to)
	if err != nil {
		return err
	}
	files = strings.Split(diff, "\n")
	if len(files) == 0 {
		return nil
	}

	// ZIPファイルを作成
	var (
		zipFile *os.File
		writer  *zip.Writer
	)
	zipFile, err = os.Create(args.ArchiveFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	writer = zip.NewWriter(zipFile)
	defer writer.Close()

	for _, file := range files {
		if file == "" {
			continue
		}

		// ファイルの内容を取得し、ZIPファイルにエントリ書き込み
		var (
			contents []byte
			entry    io.Writer
		)
		contents, err = git.Exec("show", fmt.Sprintf("%s:%s", to, file))
		if err != nil {
			return err
		}

		entry, err = writer.Create(fmt.Sprintf("%s/%s", args.ArchivePrefix, file))
		if err != nil {
			return err
		}

		if _, err = entry.Write(contents); err != nil {
			return err
		}

		if args.Verbose {
			log.Printf("ファイル追加： %s", file)
		}
	}
	if err = writer.Close(); err != nil {
		return err
	}

	return nil
}
