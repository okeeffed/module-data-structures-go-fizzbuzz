package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(i int) string {
	switch true {
	case i%3 == 0 && i%5 == 0:
		return "fizzbuzz"
	case i%3 == 0:
		return "fizz"
	case i%5 == 0:
		return "buzz"
	default:
		return strconv.Itoa(i)
	}
}

func main() {
	fmt.Println("Running fizzbuzz")
}
