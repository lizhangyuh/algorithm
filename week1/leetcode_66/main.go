package main

import "fmt"

func plusOne(digits []int) []int {
	prevPlus := 1
	for i := len(digits) - 1; i >= 0; i-- {
		current := digits[i] + prevPlus
		if current >= 10 {
			prevPlus = 1
		} else {
			prevPlus = 0
		}

		digits[i] = current % 10
	}

	if prevPlus == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {
	result := plusOne([]int{1, 2, 4})
	fmt.Printf("result: %v\n", result)
}
