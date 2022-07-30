package main

import (
	"fmt"
	"sort"
)

var ans [][]int

func main() {
	result := permuteUnique([]int{1, 2, 3})
	fmt.Printf("result: %v\n", result)
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	used := make([]bool, len(nums))
	arr := []int{}
	ans = [][]int{}

	dfs(nums, used, arr)
	return ans
}

func dfs(nums []int, used []bool, arr []int) {
	if len(arr) == len(nums) {
		ans = append(ans, append([]int{}, arr...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		arr = append(arr, nums[i])
		used[i] = true
		dfs(nums, used, arr)
		used[i] = false
		arr = arr[:len(arr)-1]
	}
}
