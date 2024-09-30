package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(coinChange([]int{2}, 3))
}

/*
Tempo: O(amount*coins)
Espaco: O(amount)
*/
func coinChange(coins []int, amount int) int {

	sort.Ints(coins)

	memo := make(map[int]int)
	memo[0] = 0

	var minCoins func(left int) int
	minCoins = func(left int) int {

		if value, ok := memo[left]; ok {
			return value
		}

		minn := math.MaxInt32

		for i := 0; i < len(coins); i++ {
			diff := left - coins[i]
			if diff < 0 {
				break
			}

			minn = min(minn, minCoins(diff)+1)
		}

		memo[left] = minn
		return minn
	}

	result := minCoins(amount)

	if result < math.MaxInt32 {
		return result
	}

	return -1

}

/*
Tempo: O(amount*coins)
Espaco: O(amount)
*/
func coinChange2(coins []int, amount int) int {

	sort.Ints(coins)

	dp := make([]int, amount+1)
	// for i := 0; i < amount+1; i++ {
	// 	dp[i] = math.MaxInt64
	// }
	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxInt32

		for _, coin := range coins {
			if coin > i {
				break
			}

			if dp[i-coin] != math.MaxInt32 {
				dp[i] = min(dp[i-coin]+1, dp[i])
			}

		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]

}
