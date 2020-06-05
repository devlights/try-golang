package enum

import (
	"github.com/devlights/gomy/output"
)

type (
	// Status -- サンプル用 enum. 何かのステータスを表す
	Status int
)

// impl -- fmt.Stringer
func (s Status) String() string {
	values := [...]string{
		"不明",
		"実行中",
		"停止中",
	}

	if s < Unknown || Stopped < s {
		return "不明"
	}

	return values[s]
}

// CanForward -- 先に進めることができるかどうかを返す
//
// 可能な場合は true, それ以外は false
func (s Status) CanForward() bool {
	switch s {
	case Running:
		return true
	}

	return false
}

// ステータス値
const (
	Unknown Status = 0 // 不明
	Running Status = 1 // 実行中
	Stopped Status = 2 // 停止中
)

// GoEnums -- Go における enum の扱い方についてのサンプルです
//
// REFERNCES:
//   -https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3
func GoEnums() error {
	// ---------------------------------------------------------------------------
	// enum とは関連する値を一つの型でグルーピングするもの
	//
	// 例： シャツのサイズ (S, M, L), ステータス (Running, Stopped, Resumed), 曜日
	// ---------------------------------------------------------------------------
	v := Running
	output.Stdoutl("[Status]", int(v))

	// ---------------------------------------------------------------------------
	// 型を定義するので、振る舞いを追加することができる
	//
	// 有名なのが golang.org/x/tools/cmd/stringer で、指定した enum の名称を出力する
	// String() を自動生成してくれるもの
	//
	// $ go get -u -v golang.org/x/tools/cmd/stringer
	// $ stringer -type Status
	//
	// 上記で status_string.go が出力される
	// ---------------------------------------------------------------------------
	output.Stdoutl("[Status]", v)
	output.Stdoutl("[v.CanForward]", v.CanForward())
	output.Stdoutl("[Running.CanForward]", Running.CanForward())
	output.Stdoutl("[Stopped.CanForward]", Stopped.CanForward())

	return nil
}
