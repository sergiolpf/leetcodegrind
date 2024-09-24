package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 2},
		//{0, 1, 1},
		//{1, 0, 1},
	}

	fmt.Println(orangesRotting(grid))
}

type position struct {
	x int
	y int
}

const (
	EMPTY  = 0
	FRESH  = 1
	ROTTEN = 2
)

/*
complexidade
Tempo: O(m*n)
Espaco: O(m*n)
*/
func orangesRotting(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])
	minutes := -1
	freshOranges := 0
	queueSize := 0

	offSetValues := []position{
		position{0, 1}, position{0, -1}, position{1, 0}, position{-1, 0},
	}

	queue := []position{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == ROTTEN {
				queue = append(queue, position{x: i, y: j})
			} else if grid[i][j] == FRESH {
				freshOranges++
			}
		}
	}

	if freshOranges == 0 {
		return 0
	}

	for len(queue) > 0 {
		queueSize = len(queue)
		minutes++

		for i := 0; i < queueSize; i++ {
			pos := queue[0]
			queue = queue[1:]

			for _, offSetPos := range offSetValues {
				r, c := pos.x+offSetPos.x, pos.y+offSetPos.y
				if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == FRESH {
					freshOranges--
					grid[r][c] = ROTTEN
					queue = append(queue, position{r, c})
				}
			}
		}
	}

	if freshOranges == 0 {
		return minutes
	}

	return -1
}
