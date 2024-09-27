package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

/*
Tempo: O(max(nums) ^ n )
espaco: O(n)
*/
func canJump(nums []int) bool {

	var helper func(index int) bool
	helper = func(index int) bool {

		if index >= len(nums)-1 {
			return true
		}

		for i := 1; i <= nums[index]; i++ {
			if helper(index + i) {
				return true
			}
		}

		return false
	}

	return helper(0)

}

/*
Tempo: O(n^2)
Espaco: O(n)
*/
func canJump2(nums []int) bool {

	seen := make(map[int]bool)
	seen[len(nums)-1] = true

	var helper func(index int) bool
	helper = func(index int) bool {

		if value, ok := seen[index]; ok {
			return value
		}

		for i := 1; i <= nums[index]; i++ {
			if helper(index + i) {
				seen[index+1] = true
				return true
			}
		}

		seen[index] = false
		return seen[index]
	}

	return helper(0)

}

/*
usando um algoritmo guloso
Tempo: O(n)
Espaco: O(1)
*/
func canJump3(nums []int) bool {

	n := len(nums)
	target := n - 1

	for i := n - 1; i >= 0; i-- {
		max_jump := nums[i]
		if i+max_jump >= target {
			target = i
		}
	}

	return target == 0

}
