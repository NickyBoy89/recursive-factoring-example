package main

import (
	"fmt"
)

func factor(number int) []int {
	var factors []int
	if number == 1 {
		return []int{1}
	}
	for i := 2; len(factors) == 0; i++ {
		// fmt.Println(i, number%i)
		if number%i == 0 {
			fmt.Printf("factor found at %v\n", i)
			factors = append(factors, i)
			factors = append(factors, factor(number/i)...)
		}
	}

	return factors
}

func main() {
	input := factor(720720) // An example of a superior highly composite number
	fmt.Printf("Result %v of size %d\n", input, len(input))
}
