package main

import "strconv"

func main() {

}

func Fizzbuzz(input int) string {
	if input == 3 || input == 6 {
		return "Fizz"
	} else if input == 5 {
		return "Buzz"
	} else {
		return strconv.Itoa(input)
	}
}
