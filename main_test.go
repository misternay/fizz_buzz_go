package main

import "testing"

func TestFizzbuzz(t *testing.T) {
	t.Run("Input 1 should be return 1", func(t *testing.T) {
		result := Fizzbuzz(1)
		expected := "1"
		if result != expected {
			t.Errorf("Input 1 should be return 1 but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 2 should be return 2", func(t *testing.T) {
		result := Fizzbuzz(2)
		expected := "2"
		if result != expected {
			t.Errorf("Input 2 should be return 2 but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 3 should be return Fizz", func(t *testing.T) {
		result := Fizzbuzz(3)
		expected := "Fizz"
		if result != expected {
			t.Errorf("Input 3 should be return Fizz but result %s not same expect %s", result, expected)
		}
	})
}
