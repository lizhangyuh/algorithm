package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	result := numIslands(grid)
	fmt.Printf("result: %v\n", result)
}

func numIslands(grid [][]byte) int {
	result := 0

	var dfs func(int, int, int)

	dfs = func(row int, col int, count int) {
		if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
			return
		}

		if string(grid[row][col]) == "0" || string(grid[row][col]) == "2" {
			return
		}

		// if count > 0 {
		//     fmt.Printf("grid: %v\n", grid)
		// }
		grid[row][col] = '2'
		result += count

		dfs(row-1, col, 0) //上
		dfs(row, col+1, 0) //右
		dfs(row+1, col, 0) //下
		dfs(row, col-1, 0) //左
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			dfs(i, j, 1)
		}
	}

	return result
}
