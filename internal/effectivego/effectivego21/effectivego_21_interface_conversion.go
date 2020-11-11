package effectivego21

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

type (
	myString struct {
		value string
	}
)

// NewMyString -- 新しい myString を生成して返します.
func NewMyString() *myString {
	return &myString{}
}

// Set -- データを設定します.
func (s *myString) Set(val string) {
	s.value = val
}

func (s *myString) String() string {
	return s.value
}

// InterfaceConversion -- Effective Go - Interface conversions and type assertions の 内容についてのサンプルです。
func InterfaceConversion() error {
	/*
		https://golang.org/doc/effective_go.html#interface_conversions

		- Type Switch は、インターフェースを別の型、もしくは別のインターフェースに変換する機能
		  - キャストではない
		  - 対象の値がインターフェースの場合のみ有効。構造体ではできない
	*/
	s1 := NewMyString()
	s1.Set("helloworld")

	// -------------------------------------------------------------------------
	// 型がインターフェースの場合は type switch が可能。
	// interface{} は、名前の通り、これもインタフェース型を表す
	// type switch は、インターフェース以外はできない。
	// つまり、以下はコンパイルエラーとなる
	// 		switch s1.(type) {
	//
	// 		}
	// -------------------------------------------------------------------------
	var val interface{} = s1
	switch v := val.(type) {
	case string:
		output.Stdoutl("string", v)
	case fmt.Stringer:
		output.Stdoutl("fmt.Stringer", v)
	}

	// -------------------------------------------------------------------------
	// 望みの型が自明な場合は以下でも良い
	// ただし、以下は間違えていた場合は panic する
	//
	// 以下のように元の値がポインタなのに、ポインタじゃない形式で変換しようとするとpanicとなる
	// 		s2 := val(myString)
	// panic: interface conversion: interface {} is *effectivego21.myString, not effectivego21.myString
	// -------------------------------------------------------------------------
	s2 := val.(*myString)
	output.Stdoutl("val.(*myString)", s2)

	// -------------------------------------------------------------------------
	// panic させないために、通常は Go の "comma, ok" イディオムを利用して
	// 変換できるかどうかを追加で受け取って、それで判断する
	// -------------------------------------------------------------------------
	s3, ok := val.(myString)
	output.Stdoutl("comma ok: val.(myString)", ok, s3)

	s4, ok := val.(*myString)
	output.Stdoutl("comma ok: val.(*myString)", ok, s4)

	return nil
}
