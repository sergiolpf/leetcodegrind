package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	var num int

	if len(nums) == 1 {
		return nums[0]
	}
	count := 0
	num = nums[0]

	for _, value := range nums {
		if value == num {
			count++
		} else {
			count--
			if count == 0 {
				count++
				num = value
			}
		}
	}

	return num
}
