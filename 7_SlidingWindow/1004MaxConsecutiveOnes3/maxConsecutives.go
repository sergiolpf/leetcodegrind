package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}

func longestOnes(nums []int, k int) int {
	left := 0
	right := 0
	maxCount := 0
	aux := 0

	if len(nums) == 1 {
		return 1
	}

	for left, right = 0, 0; right < len(nums); {
		if nums[right] == 1 {
			right++
		} else {
			if aux < k {
				aux++
				right++
			} else {
				maxCount = max(maxCount, right-left)
				for left <= right && aux == k {
					if nums[left] == 0 {
						aux--
					}
					left++
				}
			}
		}
	}

	return max(maxCount, right-left)
}
