package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	//
	// iniファイル読み込み
	//
	var (
		cfg *ini.File
		err error
	)

	cfg, err = ini.Load("config.ini")
	if err != nil {
		return err
	}

	//
	// セクション取得
	//
	var (
		secGeneral  *ini.Section
		secDatabase *ini.Section
		secLogging  *ini.Section
	)

	secGeneral = cfg.Section("General")
	secDatabase = cfg.Section("Database")
	secLogging = cfg.Section("Logging")

	//
	// キー取得
	//
	var (
		keyVersion *ini.Key
		keyHost    *ini.Key
		keyLevel   *ini.Key
	)

	keyVersion = secGeneral.Key("Version")
	keyHost = secDatabase.Key("Host")
	keyLevel = secLogging.Key("Level")

	//
	// 値取得
	//   *ini.Keyにいろいろな型で値を取得するメソッドが用意されている
	//
	fmt.Printf("Version=%s\n", keyVersion.String())
	fmt.Printf("Host=%s\n", keyHost.MustString("unknown"))
	fmt.Printf("Level=%s (%s)\n", keyLevel.Value(), keyLevel.Comment)

	return nil

	/*
	   $ task
	   task: [default] go run main.go
	   Version=1.0.0
	   Host=localhost
	   Level=DEBUG (; ログレベル)
	*/
}
