package main

import "strconv"

func main() {

}

func Fizzbuzz(input int) string {
	if input%3 == 0 && input%5 == 0 {
		return "FizzBuzz"
	} else if input%3 == 0 {
		return "Fizz"
	} else if input%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(input)
	}
}
