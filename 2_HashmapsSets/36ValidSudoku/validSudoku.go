package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	for x := 0; x < len(board); x++ {
		if !validateRows(board, x) {
			return false
		}
	}

	for y := 0; y < len(board[0]); y++ {
		if !validateCols(board, y) {
			return false
		}
	}

	m_quadrantes := make(map[string]map[byte]int)

	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if board[x][y] == '.' {
				continue
			}

			qx, qy := x/3, y/3
			key := fmt.Sprintf("%v", qx) + fmt.Sprintf("%v", qy)
			_, ok := m_quadrantes[key]

			if !ok {
				m_quadrantes[key] = make(map[byte]int, 9)
			}

			m_quadrantes[key][board[x][y]]++
			if m_quadrantes[key][board[x][y]] == 2 {
				return false
			}
		}
	}

	return true
}

func validateRows(board [][]byte, row int) bool {
	m_row := make(map[byte]int)

	for ind := 0; ind < len(board[row]); ind++ {
		if board[row][ind] != '.' {
			m_row[board[row][ind]]++
			if m_row[board[row][ind]] == 2 {
				return false
			}
		}
	}

	return true
}

func validateCols(board [][]byte, col int) bool {
	m_row := make(map[byte]int)

	for ind := 0; ind < len(board); ind++ {
		if board[ind][col] != '.' {
			m_row[board[ind][col]]++
			if m_row[board[ind][col]] == 2 {
				return false
			}
		}
	}

	return true
}
