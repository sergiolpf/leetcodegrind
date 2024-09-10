package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
}

func minSubArrayLen(target int, nums []int) int {

	if len(nums) == 1 {
		if nums[0] >= target {
			return 1
		}
		return 0
	}

	left, right := 0, 0
	minn := math.MaxInt32

	summ := 0

	for right < len(nums) {
		summ += nums[right]

		for summ >= target {
			minn = min(minn, (right - left + 1))
			summ -= nums[left]
			left++
		}

		right++
	}

	if minn != math.MaxInt32 {

		return minn
	}

	return 0
}
