package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{
		{2, 3},
		{4, 5},
		{6, 7},
		{8, 9},
		{1, 10},
	}))
}

/*
1-3, 2-6
1-6 8-10 15-18

[[1,4],[0,2],[3,5]]

1-3 3-6 10-15 2-4
1-3 2-4 3-6 10-15

1-6 10-15
*/

func merge(intervals [][]int) [][]int {

	if len(intervals) == 1 {
		return intervals
	}

	resultado := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	indResultado := -1

	for ind := 0; ind < len(intervals); ind++ {
		if len(resultado) == 0 || resultado[indResultado][1] < intervals[ind][0] {
			resultado = append(resultado, intervals[ind])
			indResultado++

		} else {
			resultado[indResultado][1] = max(intervals[ind][1], resultado[indResultado][1])

		}

	}

	return resultado

}
