package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{-8, -4, 9, 9, 4, 6, 1, -4, -1, 6, 8}))
}

func longestConsecutive(nums []int) int {
	mapaDeNums := make(map[int]bool)

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	for _, value := range nums {
		mapaDeNums[value] = true
	}

	currentCoutn := 1
	longestCount := 1

	for key := range mapaDeNums {
		if _, exists := mapaDeNums[key-1]; !exists {
			currentNum := key
			currentCoutn = 1
			for {
				if mapaDeNums[currentNum+1] {
					currentNum++
					currentCoutn++
				} else {
					break
				}

			}
			longestCount = max(currentCoutn, longestCount)
		}
	}

	return longestCount
}
