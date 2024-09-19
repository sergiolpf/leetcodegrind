package main

import "fmt"

func main() {
	//[4,3],[1,4],[4,8],[1,7],[6,4],[4,2],[7,4],[4,0],[0,9],[5,4]
	grapth := [][]int{
		{0, 1},
		{0, 2},
		{3, 5},
		{5, 4},
		{4, 3},
	}
	fmt.Println(validPath(3, grapth, 0, 5))
}

/*
Complexidade
Tempo: O(E) + O(E+V) - O(E) pq construimos o mapa.
Espaco: O(E+V)
*/
func validPath(n int, edges [][]int, source int, destination int) bool {

	if len(edges) == 0 {
		return false
	}

	mapOfEdges := make(map[int][]int)
	seen := make(map[int]struct{})

	for _, edge := range edges {
		mapOfEdges[edge[0]] = append(mapOfEdges[edge[0]], edge[1])
		mapOfEdges[edge[1]] = append(mapOfEdges[edge[1]], edge[0])
	}

	seen[source] = struct{}{}

	var dfs func(curr int) bool

	dfs = func(curr int) bool {

		if curr == destination {
			return true
		}

		if ends, ok := mapOfEdges[curr]; ok {
			for _, v := range ends {
				if _, seenn := seen[v]; !seenn {
					seen[v] = struct{}{}
					if dfs(v) {
						return true
					}
				}
			}
		}

		return false

	}

	return dfs(source)

}
