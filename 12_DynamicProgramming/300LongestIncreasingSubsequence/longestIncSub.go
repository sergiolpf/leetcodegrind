package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	//fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}

/*
Tempo: O(N^2)
Espaco: O(n)
*/
func lengthOfLIS(nums []int) int {
	size := len(nums)
	dp := make([]int, size)

	for i := 0; i < size; i++ {
		dp[i] = 1
	}

	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	maxSubSeq := math.MinInt16
	for _, num := range dp {
		maxSubSeq = max(maxSubSeq, num)
	}

	return maxSubSeq
}

/*
Tempo: O(NLogN)
Espaco: O(n)
*/
func lengthOfLIS2(nums []int) int {
	n := len(nums)

	piles := make([]int, n)
	piles[0] = nums[0]
	size := 1

	for i := 1; i < n; i++ {
		if nums[i] > piles[size-1] {
			piles[size] = nums[i]
			size++
		} else {
			j := sort.SearchInts(piles[:size], nums[i])
			piles[j] = nums[i]
		}

	}

	return size

}
