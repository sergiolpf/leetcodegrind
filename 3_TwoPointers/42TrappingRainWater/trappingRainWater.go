package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	//fmt.Println(trap([]int{4, 2, 3}))
}

func trap(height []int) int {

	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))
	maxHight, totalVol := 0, 0

	maxLeft[0] = 0
	maxRight[len(height)-1] = 0

	maxHight = height[0]
	for ind := 1; ind <= len(height)-1; ind++ {
		maxLeft[ind] = max(maxLeft[ind-1], maxHight)
		maxHight = max(maxHight, height[ind])

	}
	maxHight = height[len(height)-1]

	for ind := len(height) - 2; ind >= 0; ind-- {
		maxRight[ind] = max(maxRight[ind+1], maxHight)
		maxHight = max(maxHight, height[ind])

	}

	for ind, currentHeight := range height {
		if min(maxLeft[ind], maxRight[ind])-currentHeight > 0 {
			totalVol += min(maxLeft[ind], maxRight[ind]) - currentHeight

		}
	}

	return totalVol
}
