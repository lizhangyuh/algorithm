package main

import "fmt"

func main() {
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
	result := findRedundantDirectedConnection(edges)

	fmt.Printf("result: %v\n", result)
}

func findRedundantDirectedConnection(edges [][]int) []int {
	to := make(map[int][]int)
	degrees := make([]int, len(edges)+1)

	for _, edge := range edges {
		to[edge[0]] = append(to[edge[0]], edge[1])
		degrees[edge[1]]++
	}

	// fmt.Printf("to: %v\n", to)
	// fmt.Printf("degress: %v\n", degrees)

	valid := func() bool {
		var root int
		for i := 1; i < len(degrees); i++ {
			if degrees[i] == 0 {
				root = i
			}

			// 有入度大于0的点，表示有环
			if degrees[i] > 1 {
				return false
			}
		}

		// 没有入度为0的点，表示有环
		if root == 0 {
			return false
		}

		seen := []int{}

		var bfs func(int)

		bfs = func(root int) {
			seen = append(seen, root)
			for _, c := range to[root] {
				bfs(c)
			}
		}

		bfs(root)

		return len(seen) == len(edges)
	}

	// 倒序删除边
	for i := len(edges) - 1; i >= 0; i-- {
		edge := edges[i]
		for j, t := range to[edge[0]] {
			if t == edge[1] {
				to[edge[0]] = append(to[edge[0]][:j], to[edge[0]][j+1:]...)
				break
			}
		}
		degrees[edge[1]]--
		if valid() {
			return edges[i]
		}

		// 还原
		to[edge[0]] = append(to[edge[0]], edge[1])
		degrees[edge[1]]++
	}

	return []int{}
}
