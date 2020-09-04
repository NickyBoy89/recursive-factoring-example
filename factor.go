package main

import (
	"fmt"
)

func factor(number int) []int {
	var factors []int
	for i := 1; i <= number; i++ {
		fmt.Println(i)
		if number%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func main() {
	input := factor(720720) // An example of a superior highly composite number
	fmt.Printf("Result %v of size %d\n", input, len(input))
}
