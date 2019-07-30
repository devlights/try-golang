package constants

import "fmt"

// 定数 の 宣言
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
}
