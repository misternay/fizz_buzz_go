package main

import "strconv"

func main() {

}

func Fizzbuzz(input int) string {
	if isFizzBuzz(input) {
		return "FizzBuzz"
	} else if input%3 == 0 {
		return "Fizz"
	} else if input%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(input)
	}
}

func isFizzBuzz(input int) bool {
	return input%3 == 0 && input%5 == 0
}
