package main

import "fmt"

func main() {

	fmt.Println(findOrder(4, [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	}))
}

const (
	VISITED   = 2
	VISITING  = 1
	UNVISITED = 0
)

/*
Complexidade
Tempo: O(N + E)
Espaco: O(N + E)
*/
func findOrder(numCourses int, prerequisites [][]int) []int {

	mapPath := make(map[int][]int)

	paths := []int{}
	states := make([]int, numCourses)

	for _, prereq := range prerequisites {
		mapPath[prereq[0]] = append(mapPath[prereq[0]], prereq[1])
	}

	var dfs func(ind int) bool

	dfs = func(ind int) bool {
		state := states[ind]

		if state == VISITED {
			return true
		}

		if state == VISITING {
			return false
		}

		states[ind] = VISITING
		for _, value := range mapPath[ind] {
			if !dfs(value) {
				return false
			}
		}
		paths = append(paths, ind)
		states[ind] = VISITED

		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return []int{}
		}

	}

	return paths
}
