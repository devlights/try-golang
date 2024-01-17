package main

import "gopkg.in/ini.v1"

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
	// 既存のキーの値を変更
	//
	cfg.Section("Database").Key("User").SetValue("gitpod")

	//
	// 新しいセクションとキーの追加
	//
	newSection := cfg.Section("NewSection")
	newSection.Key("NewKey").SetValue("NewValue")

	//
	// 書き出し
	//
	if err := cfg.SaveTo("config-updated.ini"); err != nil {
		return err
	}

	return nil

	/*
	   $ task
	   task: [default] go run main.go
	   task: [default] diff -u config.ini config-updated.ini
	   --- config.ini  2024-01-17 05:18:00.670661597 +0000
	   +++ config-updated.ini  2024-01-17 05:19:09.686594430 +0000
	   @@ -5,7 +5,7 @@

	    [Database]
	    Host         = localhost
	   -User         = root
	   +User         = gitpod
	    Password     =
	    DatabaseName = my_database

	   @@ -14,3 +14,6 @@
	    Level   = DEBUG
	    ; ログファイル
	    LogFile = /var/log/myapp.log
	   +
	   +[NewSection]
	   +NewKey = NewValue
	*/

}
