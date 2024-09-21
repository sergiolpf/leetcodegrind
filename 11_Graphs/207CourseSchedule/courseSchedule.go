package main

import "fmt"

func main() {
	fmt.Println(canFinish(2, [][]int{
		{1, 0},
		{0, 1},
	}))
}

const (
	UNVISITED = 0
	VISITING  = 1
	VISITED   = 2
)

/*
Complexidade:
tempo: O(numCourses + len(prerequisites))
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	nodes := make(map[int][]int)
	states := make([]int, numCourses)

	for _, prereq := range prerequisites {
		nodes[prereq[0]] = append(nodes[prereq[0]], prereq[1])
	}

	var dfs func(node int) bool

	dfs = func(node int) bool {
		state := states[node]
		if state == VISITED {
			return true
		} else if state == VISITING {
			return false
		}

		states[node] = VISITING
		for _, num := range nodes[node] {
			if !dfs(num) {
				return false
			}
		}
		states[node] = VISITED
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}
