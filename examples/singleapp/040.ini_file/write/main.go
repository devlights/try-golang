package main

import "gopkg.in/ini.v1"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	//
	// 空のiniファイルを用意
	//
	cfg := ini.Empty()

	//
	// セクションの追加
	//
	var (
		secGeneral  *ini.Section
		secDatabase *ini.Section
		secLogger   *ini.Section
	)

	secGeneral = cfg.Section("General")
	secDatabase = cfg.Section("Database")
	secLogger = cfg.Section("Logging")

	//
	// キーと値の追加
	//
	secGeneral.NewKey("AppName", "MyApplication")
	secGeneral.NewKey("Version", "1.0.0")
	secGeneral.NewKey("License", "MIT")

	secDatabase.NewKey("Host", "localhost")
	secDatabase.NewKey("User", "root")
	secDatabase.NewKey("Password", "")
	secDatabase.NewKey("DatabaseName", "my_database")

	secLogger.NewKey("Level", "DEBUG")
	key, _ := secLogger.NewKey("LogFile", "/var/log/myapp.log")
	key.Comment = "ログファイル"

	//
	// 書き出し
	//
	if err := cfg.SaveTo("config.ini"); err != nil {
		return err
	}

	return nil

	/*
	   $ task
	   task: [default] go run main.go
	   task: [default] cat config.ini
	   [General]
	   AppName = MyApplication
	   Version = 1.0.0
	   License = MIT

	   [Database]
	   Host         = localhost
	   User         = root
	   Password     =
	   DatabaseName = my_database

	   [Logging]
	   Level   = DEBUG
	   ; ログファイル
	   LogFile = /var/log/myapp.log
	*/
}
