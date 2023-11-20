package constants

import "fmt"

// Pi -- 定数 の 宣言
const Pi = 3.14

// 定数 を 一気に宣言
const (
	Ng      = 0
	Ok      = 1
	Unknown = 99
)

// ConstStatementDeclares -- const による 定数 の宣言についてのサンプル
func ConstStatementDeclares() error {

	fmt.Printf("Pi=%f, (%d, %d, %d)\n", Pi, Ng, Ok, Unknown)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: const_statement_declare

	   [Name] "const_statement_declare"
	   Pi=3.140000, (0, 1, 99)


	   [Elapsed] 7.65µs
	*/

}
