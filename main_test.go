package main

import "testing"

func TestFizzbuzz(t *testing.T) {
	t.Run("Input 1 should be return 1", func(t *testing.T) {
		result := Fizzbuzz(1)
		expected := "1"
		if result != expected {
			t.Errorf("Input 1 should be return 1", result, expected)
		}
	})
}
