package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(i int) (s string) {
	if i%3 == 0 && i%5 == 0 {
		s = "FizzBuzz"
	} else if i%3 == 0 {
		s = "Fizz"
	} else if i%5 == 0 {
		s = "Buzz"
	} else {
		s = strconv.Itoa(i)
	}
	return
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(FizzBuzz(i))
	}
}