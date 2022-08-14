package main

import "fmt"

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5

	result := shipWithinDays(weights, days)

	fmt.Printf("result: %v\n", result)
}

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	for _, w := range weights {
		right += w
		if left < w {
			left = w
		}
	}

	var day, sum int
	result := right

	for left <= right {
		day = 1
		sum = 0
		mid := (left + right) / 2
		// fmt.Printf("mid: %v\n", mid)

		for _, w := range weights {
			if sum+w > mid {
				day++
				sum = 0
			}
			sum += w
		}

		// fmt.Printf("day: %v\n", day)

		if day <= days {
			if mid < result {
				result = mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}
