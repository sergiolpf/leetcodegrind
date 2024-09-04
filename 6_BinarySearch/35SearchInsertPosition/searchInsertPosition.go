package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3}, 0))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		} else {
			return 1
		}
	}

	left := 0
	right := len(nums) - 1
	middle := 0

	for left <= right {
		middle = left + (right-left)/2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}

	}

	return left

}
