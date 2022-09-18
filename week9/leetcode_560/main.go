package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	k := 2

	result := subarraySum(nums, k)
	fmt.Printf("result: %v\n", result)
}

func subarraySum(nums []int, k int) int {
	pre, ans := 0, 0
	m := map[int]int{}
	m[0] = 1

	for i := 0; i < len(nums); i++ {
		pre += nums[i]

		if _, ok := m[pre-k]; ok {
			ans += m[pre-k]
		}
		m[pre] += 1
	}

	return ans
}
