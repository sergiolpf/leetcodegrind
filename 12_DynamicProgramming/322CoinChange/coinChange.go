package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(coinChange([]int{2}, 3))
}

func coinChange(coins []int, amount int) int {

	sort.Ints(coins)

	memo := make(map[int]int)
	memo[0] = 0

	var minCoins func(left int) int
	minCoins = func(left int) int {

		if value, ok := memo[left]; ok {
			return value
		}

		minn := math.MaxInt64

		for i := 0; i < len(coins); i++ {
			diff := left - coins[i]
			if diff < 0 {
				break
			}

			minCoinDiff := minCoins(diff)
			if minCoinDiff == math.MaxInt64 {
				minn = min(minn, minCoinDiff)
			} else {
				minn = min(minn, minCoinDiff+1)
			}
		}

		memo[left] = minn
		return minn
	}

	result := minCoins(amount)

	if result < math.MaxInt64 {
		return result
	}

	return -1

}
