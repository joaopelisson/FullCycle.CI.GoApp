package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(10, 5)
	expected := 15

	if result != expected {
		t.Errorf("Sum(10, 5) = %d; want %d", result, expected)
	}
}