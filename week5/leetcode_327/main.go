package main

import "fmt"

func main() {
	result := countRangeSum([]int{-2, 5, -1}, -2, 2)
	fmt.Printf("result: %v\n", result)
}

func countRangeSum(nums []int, lower int, upper int) int {
	// 计算前缀和
	preSum := make([]int, len(nums)+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}

	var countSum func([]int) int

	countSum = func(arr []int) int {
		length := len(arr)
		if length <= 1 {
			return 0
		}

		left := append([]int{}, arr[:length/2]...)
		right := append([]int{}, arr[length/2:]...)

		count := countSum(left) + countSum(right)

		l, r := 0, 0
		for _, v := range left {
			for l < len(right) && right[l]-v < lower {
				l++
			}

			for r < len(right) && right[r]-v <= upper {
				r++
			}

			count += r - l
		}

		// 合并左右数组
		p1, p2 := 0, 0
		for i := range arr {
			if p1 < len(left) && (p2 == len(right) || left[p1] <= right[p2]) {
				arr[i] = left[p1]
				p1++
			} else {
				arr[i] = right[p2]
				p2++
			}
		}

		return count
	}

	return countSum(preSum)
}
