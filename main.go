package main

import "strconv"

func main() {

}

func Fizzbuzz(input int) string {
	if isFizzBuzz(input) {
		return "FizzBuzz"
	} else if isFizz(input) {
		return "Fizz"
	} else if isBuzz(input) {
		return "Buzz"
	} else {
		return strconv.Itoa(input)
	}
}

func isFizzBuzz(input int) bool {
	return input%3 == 0 && input%5 == 0
}

func isFizz(input int) bool {
	return input%3 == 0
}
func isBuzz(input int) bool {
	return input%5 == 0
}
