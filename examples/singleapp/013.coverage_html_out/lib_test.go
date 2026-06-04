package main

import "testing"

func TestPlus(t *testing.T) {
	// Arrange
	var (
		x        = 1
		y        = 2
		expected = 3
	)

	// Act
	result := Plus(x, y)

	// Assert
	if result != expected {
		t.Errorf("[want] %v\t[got] %v", expected, result)
	}
}
