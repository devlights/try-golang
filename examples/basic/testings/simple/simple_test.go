package simple

import (
	"math"
	"testing"
)

func TestAddSimpleSuccess(t *testing.T) {
	// Arrange
	values := []int32{1, 2, 3, 4, 5}
	want := int32(15)

	// Act
	got, err := Add(values...)

	// Assert
	if err != nil {
		t.Errorf("error returns %v", err)
	}

	if want != got {
		t.Errorf("Add() = %v, want %v", got, want)
	}
}

func TestAddSimpleOverflow(t *testing.T) {
	// Arrange
	values := []int32{1, math.MaxInt32}
	want := int32(1)

	// Act
	got, err := Add(values...)

	// Assert
	if err == nil {
		t.Error("not error")
	}

	if err != ErrOverflow {
		t.Errorf("illegal error type %t", err)
	}

	if got != want {
		t.Errorf("Add() = %v, want %v", got, want)
	}
}

func TestAddSimpleUnderflow(t *testing.T) {
	// Arrange
	values := []int32{-1, math.MinInt32}
	want := int32(-1)

	// Act
	got, err := Add(values...)

	// Assert
	if err == nil {
		t.Error("not error")
	}

	if err != ErrUnderflow {
		t.Errorf("illegal error type %t", err)
	}

	if got != want {
		t.Errorf("Add() = %v, want %v", got, want)
	}
}
