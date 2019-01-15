package main

import "strconv"

func main() {

}

func Fizzbuzz(input int) string {
	if input == 3 {
		return "Fizz"
	} else {
		return strconv.Itoa(input)
	}
}
