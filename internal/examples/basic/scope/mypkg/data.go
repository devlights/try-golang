package mypkg

import "fmt"

// ICanDisplayValues -- サンプル用インターフェース
type ICanDisplayValues interface {
	// GetValues -- サンプル用インターフェースメソッド
	GetValues() string
}

// HasPublicFields -- サンプル用の構造体
// public な フィールドを持つ
type HasPublicFields struct {
	// Val1 -- 値１
	Val1 int
	// Val2 -- 値２
	Val2 string
}

// HasNoPublicFields -- サンプル用の構造体
// private な フィールドを持つ
type HasNoPublicFields struct {
	val1 int
	val2 string
}

// GetValues -- サンプル用メソッド
func (h *HasPublicFields) GetValues() string {
	return fmt.Sprintf("%+v", h)
}

// GetValues -- サンプル用メソッド
func (h *HasNoPublicFields) GetValues() string {
	return fmt.Sprintf("%+v", h)
}
