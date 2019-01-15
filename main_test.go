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
	t.Run("Input 10 should be return Buzz", func(t *testing.T) {
		result := Fizzbuzz(10)
		expected := "Buzz"
		if result != expected {
			t.Errorf("Input 10 should be return Buzz but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 11 should be return 11", func(t *testing.T) {
		result := Fizzbuzz(11)
		expected := "11"
		if result != expected {
			t.Errorf("Input 11 should be return 11 but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 12 should be return Fizz", func(t *testing.T) {
		result := Fizzbuzz(12)
		expected := "Fizz"
		if result != expected {
			t.Errorf("Input 12 should be return Fizz but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 13 should be return 13", func(t *testing.T) {
		result := Fizzbuzz(13)
		expected := "13"
		if result != expected {
			t.Errorf("Input 13 should be return 13 but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 14 should be return 14", func(t *testing.T) {
		result := Fizzbuzz(14)
		expected := "14"
		if result != expected {
			t.Errorf("Input 14 should be return 14 but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 15 should be return FizzBuzz", func(t *testing.T) {
		result := Fizzbuzz(15)
		expected := "FizzBuzz"
		if result != expected {
			t.Errorf("Input 15 should be return FizzBuzz but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 16 should be return 16", func(t *testing.T) {
		result := Fizzbuzz(16)
		expected := "16"
		if result != expected {
			t.Errorf("Input 16 should be return 16 but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 17 should be return 17", func(t *testing.T) {
		result := Fizzbuzz(17)
		expected := "17"
		if result != expected {
			t.Errorf("Input 17 should be return 17 but result %s not same expect %s", result, expected)
		}
	})
	t.Run("Input 18 should be return 18", func(t *testing.T) {
		result := Fizzbuzz(18)
		expected := "Fizz"
		if result != expected {
			t.Errorf("Input 18 should be return Fizz but result %s not same expect %s", result, expected)
		}
	})
}
