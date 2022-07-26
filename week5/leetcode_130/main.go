package main

import "fmt"

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)

	fmt.Printf("board: %v\n", board)

}

func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	q := [][]int{}

	var bfs func(int, int)

	bfs = func(x int, y int) {
		q = q[1:]

		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]

			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && board[nx][ny] == 'O' {
				q = append(q, []int{nx, ny})
				board[nx][ny] = '#'
			}
		}
	}

	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			q = append(q, []int{i, 0})
			board[i][0] = '#'
		}
		if board[i][cols-1] == 'O' {
			q = append(q, []int{i, cols - 1})
			board[i][cols-1] = '#'
		}
	}

	for i := 1; i < cols-1; i++ {
		if board[0][i] == 'O' {
			q = append(q, []int{0, i})
			board[0][i] = '#'
		}
		if board[rows-1][i] == 'O' {
			q = append(q, []int{rows - 1, i})
			board[rows-1][i] = '#'
		}
	}

	// fmt.Printf("board: %v\n", board)

	for len(q) > 0 {
		bfs(q[0][0], q[0][1])
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}
