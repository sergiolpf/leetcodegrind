package main

import "fmt"

func main() {

	fmt.Println(search([]int{2, 5}, 5))

}

func search(nums []int, target int) int {

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2

		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}
