package main

import "fmt"

func main() {
	result := longestPalindromeSubseq("bbbab")
	fmt.Printf("result: %v\n", result)
}

func longestPalindromeSubseq(s string) int {
	// s[i] != s[j]
	// dp[i][j] = max(dp[i][j-1], dp[i+1][j])
	// s[i] == s[j]
	// dp[i][j] = dp[i+1][j-1] + 2

	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
			// fmt.Printf("dp[%v][%v], %v\n", i, j, dp[i][j])
		}
	}

	return dp[0][len(s)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
