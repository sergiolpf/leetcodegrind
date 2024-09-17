package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	response, solution := [][]int{}, []int{}

	if len(nums) == 1 {
		return append(response, nums)
	}

	backtrack(&response, solution, nums, map[int]struct{}{})

	return response
}

/*
Complexidad
Tempo: O(n2)
Espaco: O(n)
*/
func backtrack(res *[][]int, sol []int, nums []int, used map[int]struct{}) {
	if len(sol) == len(nums) {
		temp := make([]int, len(sol))
		copy(temp, sol)
		*res = append(*res, temp)
		return
	}

	for _, num := range nums {
		if _, ok := used[num]; !ok {
			used[num] = struct{}{}
			sol = append(sol, num)

			backtrack(res, sol, nums, used)

			delete(used, num)
			sol = sol[:len(sol)-1]
		}
	}
}
