package main

import "testing"

// SampleMapping.MakeMappings の テスト
func TestMakeMapping(t *testing.T) {
	// Arrange
	sut := make(SampleMapping)

	// Act
	sut.MakeMapping()

	// Assert
	if len(sut) == 0 {
		t.Errorf("[NG] Mapping Count=0")
	}
}

// サンプルが取得できるかどうか の テスト
func TestRetriveExample_Success(t *testing.T) {
	// Arrange
	sut := make(SampleMapping)

	// Act
	sut.MakeMapping()

	// Assert
	if sut["helloworld"] == nil {
		t.Errorf("[NG] Example object is nil")
	}
}
