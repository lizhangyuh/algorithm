package main

import "fmt"

func main() {
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}

	result := findRedundantConnection(edges)
	fmt.Printf("result: %v\n", result)
}

func findRedundantConnection(edges [][]int) []int {
	group := make([]int, len(edges)+1)
	for i := 0; i < len(edges)+1; i++ {
		group[i] = i
	}

	var find func(i int) int

	// 查找真正的联通分量
	find = func(i int) int {
		if group[i] == i {
			return i
		}

		return find(group[i])
	}

	// 合并两个联通分量，返回这个两个联通分量是否本来就联通了
	union := func(from int, to int) bool {
		from = find(from)
		to = find(to)

		if from == to {
			return true
		}

		group[to] = from
		fmt.Printf("group: %v\n", group)
		return false
	}

	for _, edge := range edges {
		// 联通分量事先已经联通则表示有环
		if union(edge[0], edge[1]) {
			return edge
		}
	}

	return []int{}
}
