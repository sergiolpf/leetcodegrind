package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{1, 13, 15, 17}))
}

func findMin(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)-1

	if nums[right] > nums[left] {
		return nums[left]
	}

	for left <= right {
		pivot := left + (right-left)/2

		if nums[pivot] > nums[pivot+1] {
			return nums[pivot+1]
		}

		if pivot > 0 && nums[pivot-1] > nums[pivot] {
			return nums[pivot]
		}

		if nums[pivot] > nums[0] {
			left = pivot + 1
		} else {
			right = pivot - 1
		}

	}

	return -1
}
