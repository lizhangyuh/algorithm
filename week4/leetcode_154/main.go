package main

import "fmt"

func main() {
	result := findMin([]int{0, 1, 4, 4, 5, 6, 7})
	fmt.Printf("result: %v\n", result)
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		// 重复元素则去掉最右端的那个
		if nums[mid] == nums[right] {
			right--
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[right]
}
