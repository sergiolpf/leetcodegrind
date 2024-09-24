package main

import (
	"fmt"
)

func main() {

	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	fmt.Println(pacificAtlantic(heights))
}

/*
Complexidade:
Tempo: O(m*n)
Espaco: O(m*n)
*/
func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 1 && len(heights[0]) == 1 {
		return [][]int{{0, 0}}
	}

	results := [][]int{}

	m := len(heights)
	n := len(heights[0])
	pacificVisited := make([][]bool, m)
	atlanticVisited := make([][]bool, m)

	for i := 0; i < m; i++ {
		pacificVisited[i] = make([]bool, n)
		atlanticVisited[i] = make([]bool, n)
	}

	var dfs func(visited [][]bool, i, j, value int)
	dfs = func(visited [][]bool, i, j, value int) {

		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}

		if visited[i][j] {
			return
		}

		if heights[i][j] < value {
			return
		}

		visited[i][j] = true
		dfs(visited, i, j+1, heights[i][j])
		dfs(visited, i, j-1, heights[i][j])
		dfs(visited, i+1, j, heights[i][j])
		dfs(visited, i-1, j, heights[i][j])

	}

	for a := 0; a < m; a++ {
		dfs(pacificVisited, a, 0, -1)
		dfs(atlanticVisited, a, n-1, -1)
	}

	for a := 0; a < n; a++ {
		dfs(pacificVisited, 0, a, -1)
		dfs(atlanticVisited, m-1, a, -1)
	}

	for a := 0; a < m; a++ {
		for b := 0; b < n; b++ {
			if pacificVisited[a][b] && atlanticVisited[a][b] {
				results = append(results, []int{a, b})
			}

		}
	}
	return results

}
