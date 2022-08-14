package main

import "fmt"

func main() {
	result := findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	fmt.Printf("result: %v\n", result)
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	result := []int{}
	pre := make(map[int][]int)
	post := make(map[int][]int)

	for _, req := range prerequisites {
		x := req[0]
		y := req[1]
		pre[x] = append(pre[x], y)
		post[y] = append(post[y], x)
	}

	// fmt.Printf("pre: %v\n", pre)
	// fmt.Printf("post: %v\n", post)

	queue := []int{}

	for i := 0; i < numCourses; i++ {
		if _, ok := pre[i]; !ok {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		result = append(result, cur)
		// fmt.Printf("queue: %v\n", queue)
		queue = queue[1:]
		for _, p := range post[cur] {
			for i, q := range pre[p] {
				if q == cur {
					pre[p] = append(pre[p][:i], pre[p][i+1:]...)
					if len(pre[p]) == 0 {
						queue = append(queue, p)
					}
					break
				}
			}
		}

	}

	if len(result) != numCourses {
		return []int{}
	}

	return result
}
