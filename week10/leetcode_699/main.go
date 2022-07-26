package main

import "fmt"

type Box struct {
	Left   int
	Right  int
	Height int
}

func main() {
	input := [][]int{{1, 2}, {2, 3}, {6, 1}}
	result := fallingSquares(input)
	fmt.Printf("result: %v\n", result)
}

func fallingSquares(positions [][]int) []int {
	boxes := []Box{}
	ans := []int{}
	max := 0

	for _, p := range positions {
		left := p[0]
		right := p[0] + p[1]
		height := p[1]
		bottom := 0

		for _, b := range boxes {
			if !(left >= b.Right || right <= b.Left) && bottom < b.Height {
				bottom = b.Height
			}
		}

		height += bottom
		if height > max {
			ans = append(ans, height)
			max = height
		} else {
			ans = append(ans, max)
		}

		boxes = append(boxes, Box{Left: left, Right: right, Height: height})
	}

	return ans
}
