package main

import "fmt"

func main() {

	fmt.Println(search([]int{5, 1, 3}, 3))
}

func search(nums []int, target int) int {

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	indMin := findMin(nums)
	left, right := 0, len(nums)-1

	if nums[indMin] == target {
		return indMin
	}

	if target > nums[indMin] && target <= nums[right] {
		left = indMin + 1
	} else {
		right = indMin - 1

	}

	for left <= right {
		pivot := left + (right-left)/2

		if nums[pivot] == target {
			return pivot
		}

		if nums[pivot] > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return -1
}

func findMin(nums []int) int {

	if len(nums) == 1 {
		return 0
	}

	left, right := 0, len(nums)-1

	if nums[right] > nums[left] {
		return left
	}

	for left <= right {
		pivot := left + (right-left)/2

		if nums[pivot] > nums[pivot+1] {
			return pivot + 1
		}

		if pivot > 0 && nums[pivot-1] > nums[pivot] {
			return pivot
		}

		if nums[pivot] > nums[0] {
			left = pivot + 1
		} else {
			right = pivot - 1
		}

	}

	return -1
}
