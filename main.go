package main

import "strconv"

func main() {

}

func Fizzbuzz(input int) string {
	if input == 3 {
		return "Fizz"
	} else if input == 5 {
		return "Buzz"
	} else {
		return strconv.Itoa(input)
	}
}
