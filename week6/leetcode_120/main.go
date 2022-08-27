package main

import (
	"fmt"
	"math"
)

func main() {
	result := minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	})

	fmt.Printf("result: %v\n", result)
}

func minimumTotal(triangle [][]int) int {
	// dp[i][j] = triangle[i][j] + min(dp[i-1][j], dp[i-1][j-1])

	n := len(triangle)

	if n == 1 {
		return triangle[0][0]
	}

	dp := make([]int, n)
	dp[0] = triangle[0][0]
	ans := math.MaxInt

	for i := 1; i < n; i++ {
		for j := i; j >= 0; j-- {
			if j == 0 {
				dp[j] += triangle[i][j]
			} else if j == i {
				dp[j] = triangle[i][i] + dp[j-1]
			} else if dp[j] < dp[j-1] {
				dp[j] = triangle[i][j] + dp[j]
			} else {
				dp[j] = triangle[i][j] + dp[j-1]
			}

			if i == n-1 {
				if dp[j] < ans {
					ans = dp[j]
				}
			}

		}
	}

	return ans
}
