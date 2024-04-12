package varstatement

import "fmt"

// Basic -- var による 変数 の宣言についてのサンプル
func Basic() error {
	// var により 変数を宣言することが可能
	// Go は、型の指定を後ろに配置するタイプ
	var i = 0
	var s = "hello"

	// 初期値を設定している場合、型は省略可能
	var i2 = 1
	var s2 = "world"

	// 複数の var は、以下のようにまとめることが可能
	var (
		i3 = 2
		s3 = "helloworld"
	)

	var (
		name, location string
		age            int
	)

	name = "hehe"
	location = "japan"
	age = 33

	// 複数の場合も以下のように初期化することも可能
	var (
		name2, location2, age2 = "hoge", "japan", 44
	)

	fmt.Printf("%#v,%#v,%#v,%#v,%#v,%#v\n", i, s, i2, s2, i3, s3)
	fmt.Printf("%#v,%#v,%#v,%#v,%#v,%#v\n", name, location, age, name2, location2, age2)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: var_statement_declare

	   [Name] "var_statement_declare"
	   0,"hello",1,"world",2,"helloworld"
	   "hehe","japan",33,"hoge","japan",44


	   [Elapsed] 8.809µs
	*/

}
