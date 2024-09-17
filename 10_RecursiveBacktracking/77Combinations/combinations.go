package main

import (
	"fmt"
)

func main() {

	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {

	if n == 1 {
		return [][]int{{1}}
	}

	response, solution := [][]int{}, []int{}

	backtrack(n, &response, solution, n, k)
	return response
}

/*
Complexity
Tempo: O(N choose K)
Espaco: O(N)
*/
func backtrack(ind int, res *[][]int, sol []int, n, k int) {
	if len(sol) == k {
		temp := make([]int, k)
		copy(temp, sol)
		*res = append(*res, temp)
		return
	}

	left := ind
	still_need := k - len(sol)

	if left > still_need {
		backtrack(ind-1, res, sol, n, k)
	}

	sol = append(sol, ind)

	backtrack(ind-1, res, sol, n, k)
	sol = sol[:len(sol)-1]
}
