package main

import "fmt"

func main() {
	board := [][]byte{
		{'A'},
	}

	fmt.Println(exist(board, "AB"))
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	w := len(word)

	if m == 1 && n == 1 {
		return string(board[0][0]) == word
	}

	offSet := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var backtrack func(i, j int, index int) bool

	backtrack = func(i, j int, index int) bool {

		if index == w {
			return true
		}

		if board[i][j] != ([]byte(word))[index] {
			return false
		}

		temp := board[i][j]
		board[i][j] = '#'

		for _, offSetValues := range offSet {
			r, c := i+offSetValues[0], j+offSetValues[1]

			if (0 <= r && r < m) && (0 <= c && c < n) {
				if backtrack(r, c, index+1) {
					return true
				}
			}
		}

		board[i][j] = temp

		return false
	}

	for i, line := range board {
		for j, _ := range line {
			if backtrack(i, j, 0) {
				return true
			}
		}

	}

	return false
}
