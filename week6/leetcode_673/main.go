package main

import "fmt"

func main() {
	result := findNumberOfLIS([]int{1, 3, 5, 4, 7})
	fmt.Printf("result: %v\n", result)
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	count := make([]int, n)
	maxLength := 0
	ans := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				}
			}

		}

		if dp[i] == maxLength {
			ans += count[i]
		}
		if dp[i] > maxLength {
			maxLength = dp[i]
			ans = count[i]
		}
	}

	return ans
}
