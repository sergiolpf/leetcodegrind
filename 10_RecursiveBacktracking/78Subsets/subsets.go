package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	response, solution := [][]int{}, []int{}

	if nums == nil || len(nums) == 0 {
		return response
	}

	backtrack(0, &response, solution, nums)
	return response
}

func backtrack(i int, res *[][]int, sol []int, nums []int) {

	if i == len(nums) {
		temp := make([]int, len(sol))
		copy(temp, sol)
		*res = append(*res, temp)
	}

	//Don't pick nums[i]
	backtrack(i+1, res, sol, nums)

	// Pick nums[i]
	sol = append(sol, nums[i])
	backtrack(i+1, res, sol, nums)
	sol = sol[:len(sol)-1]

}
