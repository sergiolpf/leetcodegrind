package main

import "fmt"

func main() {
	//fmt.Println(rob3([]int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}))
	fmt.Println(rob3([]int{2, 1, 1, 2}))
}

/*
Tempo: O(n^2)
Espaco: O(n)
*/
func rob(nums []int) int {

	n := len(nums)

	var helper func(ind int) int
	helper = func(ind int) int {

		if ind == 0 {
			return nums[0]
		}
		if ind == 1 {
			return max(nums[0], nums[1])
		}

		return max(nums[ind]+helper(ind-2), helper(ind-1))
	}

	return helper(n - 1)
}

/*
Tempo: O(n)
Espaco: O(n)
*/
func rob2(nums []int) int {

	n := len(nums)
	memo := make(map[int]int)

	var helper func(ind int) int
	helper = func(ind int) int {

		if _, ok := memo[ind]; ok {
			return memo[ind]
		}

		if ind == 0 {
			return nums[0]
		}
		if ind == 1 {
			return max(nums[0], nums[1])
		}

		memo[ind] = max(nums[ind]+helper(ind-2), helper(ind-1))
		return memo[ind]
	}

	return helper(n - 1)
}

/*
Tempo: O(n)
Espaco: O(n)
*/
func rob3(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	n := len(nums)
	memo := make([]int, n)
	memo[0] = nums[0]
	memo[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		memo[i] = max(nums[i]+memo[i-2], memo[i-1])
	}

	return memo[n-1]

}

/*
Tempo: O(n)
Espaco: O(1)
*/
func rob4(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	n := len(nums)
	prev, curr := nums[0], max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		prev, curr = curr, max(nums[i]+prev, curr)
	}

	return curr

}
