package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2}, 1))
}

func combinationSum(candidates []int, target int) [][]int {
	response, solution := [][]int{}, []int{}

	backtrack(0, &response, solution, candidates, 0, target)
	return response
}

func backtrack(ind int, res *[][]int, sol []int, nums []int, curr_sum int, target int) {

	if curr_sum == target {
		temp := make([]int, len(sol))
		copy(temp, sol)
		*res = append(*res, temp)
		return
	}

	if curr_sum > target || ind == len(nums) {
		return
	}

	backtrack(ind+1, res, sol, nums, curr_sum, target)

	sol = append(sol, nums[ind])
	backtrack(ind, res, sol, nums, curr_sum+nums[ind], target)
	sol = sol[:len(sol)-1]
}
