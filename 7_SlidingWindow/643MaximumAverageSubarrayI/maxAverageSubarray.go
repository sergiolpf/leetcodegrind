package main

import "fmt"

func main() {
	//fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
}

func findMaxAverage(nums []int, k int) float64 {
	left := 0
	right := k - 1
	summ := 0
	sett := make(map[int]int, len(nums))

	for aux := left; aux <= right; aux++ {
		summ += nums[aux]
		sett[nums[aux]]++
	}

	maxAvg := float64(summ) / float64(k)

	for right++; right < len(nums); right++ {
		summ -= nums[left]
		left++
		summ += nums[right]
		maxAvg = max(maxAvg, float64(summ)/float64(k))
	}

	return maxAvg

}
