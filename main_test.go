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
	t.Run("Input 4 should be return 4", func(t *testing.T) {
		result := Fizzbuzz(4)
		expected := "4"
		if result != expected {
			t.Errorf("Input 4 should be return 4 but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 5 should be return Buzz", func(t *testing.T) {
		result := Fizzbuzz(5)
		expected := "Buzz"
		if result != expected {
			t.Errorf("Input 5 should be return Buzz but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 6 should be return Fizz", func(t *testing.T) {
		result := Fizzbuzz(6)
		expected := "Fizz"
		if result != expected {
			t.Errorf("Input 6 should be return Fizz but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 7 should be return 7", func(t *testing.T) {
		result := Fizzbuzz(7)
		expected := "7"
		if result != expected {
			t.Errorf("Input 7 should be return 7 but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 8 should be return 8", func(t *testing.T) {
		result := Fizzbuzz(8)
		expected := "8"
		if result != expected {
			t.Errorf("Input 8 should be return 8 but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 9 should be return Fizz", func(t *testing.T) {
		result := Fizzbuzz(9)
		expected := "Fizz"
		if result != expected {
			t.Errorf("Input 9 should be return Fizz but result %s not same expect %s", result, expected)
		}
	})
}
