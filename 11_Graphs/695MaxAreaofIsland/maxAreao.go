package main

import "fmt"

func main() {

	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	fmt.Println(maxAreaOfIsland(grid))
}

/*
Complexidade
Tempo: O(m * n)
Espaco: O(m * n)
*/
func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	maxArea := 0
	currArea := 0

	offSet := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var dfs func(i, j int)

	dfs = func(i, j int) {

		if grid[i][j] != 1 {
			return
		}

		currArea++
		grid[i][j] = '0'

		for _, offSetValues := range offSet {
			r, c := i+offSetValues[0], j+offSetValues[1]
			if (0 <= r && r < m) && (0 <= c && c < n) {
				dfs(r, c)
			}
		}

	}

	for a, line := range grid {
		for b := range line {
			if grid[a][b] == 1 {
				dfs(a, b)
				maxArea = max(maxArea, currArea)
				currArea = 0
			}
		}
	}

	return maxArea
}
