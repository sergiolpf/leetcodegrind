package main

import "fmt"

func main() {

	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'1', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Println(numIslands(grid))
}

/*
Complexidade
Tempo: O(n*m)
espaco: O(n*m)
*/
func numIslands(grid [][]byte) int {

	m := len(grid)
	n := len(grid[0])
	count := 0

	offSet := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var dfs func(i, j int)
	dfs = func(i, j int) {

		if (i < 0 || i >= m) || (j < 0 || j >= n) || grid[i][j] != '1' {
			return
		}

		grid[i][j] = '0'

		for _, offSetValues := range offSet {
			r, c := i+offSetValues[0], j+offSetValues[1]

			{
				dfs(r, c)
			}
		}

	}

	for i, line := range grid {
		for j := range line {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}

	}

	return count
}
