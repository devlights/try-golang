package lib

import (
	"testing"

	"github.com/devlights/try-golang/basic/helloworld"
	"github.com/devlights/try-golang/interfaces"
)

type (
	testExampleRegister struct{}
)

func (t *testExampleRegister) Regist(m interfaces.ExampleMapping) {
	m["helloworld"] = func() error {
		return nil
	}
}

// SampleMapping.MakeMappings の テスト
func TestMakeMapping(t *testing.T) {
	// Arrange
	register := new(testExampleRegister)
	sut := make(interfaces.ExampleMapping)

	// Act
	sut.MakeMapping(register)

	// Assert
	if len(sut) == 0 {
		t.Errorf("[NG] Mapping Count=0")
	}
}

// サンプルが取得できるかどうか の テスト
func TestRetriveExample_Success(t *testing.T) {
	// Arrange
	register := new(testExampleRegister)
	sut := make(interfaces.ExampleMapping)

	// Act
	sut.MakeMapping(register)

	// Assert
	if sut["helloworld"] == nil {
		t.Errorf("[NG] Example object is nil")
	}
}

// ExampleHelloworld -- helloworld サンプルの実行
//
// golang には、「Examples」 という機能がある.
// https://golang.org/pkg/testing/#hdr-Examples
// コメントで Output: という行を書いて、その行またはその下に
// 実行された結果を記載することでユニットテスト時に検証してくれる.
// 実装した関数の動作サンプルとして記載するのにちょうど良い。
// 出力結果の比較は、「標準出力」に出た結果と比較される。
//
// 命名規則は、Exampleで始める名前となっていること。
func ExampleHelloWorld() {

	_ = helloworld.HelloWorld()

	// Output: Hello World!
}
