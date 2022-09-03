package main

import (
	"fmt"
	"math"
)

func main() {
	result := numSquares(12)
	fmt.Printf("result: %v\n", result)
}

func numSquares(n int) int {
	// dp[i] = min(dp[i-j^2]) + 1

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minDp := math.MaxInt
		for j := 1; j*j <= i; j++ {
			minDp = min(minDp, dp[i-j*j])
		}
		dp[i] = minDp + 1
	}

	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
