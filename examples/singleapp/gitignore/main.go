package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/integrii/flaggy"
)

const (
	BASE_URL = "https://gitignore.io/api/"
)

var (
	langCmd *flaggy.Subcommand
	listCmd *flaggy.Subcommand

	lang string
)

func init() {
	langCmd = flaggy.NewSubcommand("lang")
	langCmd.Description = "指定した言語で .gitignore を出力"
	langCmd.AddPositionalValue(&lang, "", 1, true, "言語 (ex: go, python, java,,,)")

	listCmd = flaggy.NewSubcommand("list")
	listCmd.Description = "利用可能言語リストを出力"

	flaggy.AttachSubcommand(langCmd, 1)
	flaggy.AttachSubcommand(listCmd, 1)

	flaggy.SetName("gitignore")
	flaggy.SetDescription(".gitignoreを生成するツール")
	flaggy.SetVersion("v1.0.0")
}

func main() {
	flaggy.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		err error
	)

	switch {
	case langCmd.Used:
		err = request(fmt.Sprintf("%s%s", BASE_URL, lang))
	case listCmd.Used:
		err = request(fmt.Sprintf("%s%s", BASE_URL, "list"))
	default:
		flaggy.ShowHelp("")
	}

	if err != nil {
		return err
	}

	return nil
}

func request(url string) error {
	var (
		res *http.Response
		err error
	)

	res, err = http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		return err
	}

	return nil
}
