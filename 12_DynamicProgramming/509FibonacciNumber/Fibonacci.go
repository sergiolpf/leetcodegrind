package main

import "fmt"

func main() {
	fmt.Println(fibButtonUp(8))
}

// recursive
func fibRecursive(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibRecursive(n-1) + fibRecursive(n-2)
}

/*
complexity
Tempo: O(n)
Espaco: O(n)
*/
func fibMemo(n int) int {
	memo := make(map[int]int)
	memo[0] = 0
	memo[1] = 1

	var memoization func(x int) int
	memoization = func(x int) int {

		if _, ok := memo[x]; ok {
			return memo[x]
		}

		memo[x] = memoization(x-2) + memoization(x-1)
		return memo[x]
	}

	return memoization(n)
}

/*
Complexidade:
Tempo: O(n)
Espaco: O(n)
*/
func fibButtonUp(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	table := make([]int, n+1)

	table[1] = 1

	for i := 2; i < n+1; i++ {
		table[i] = table[i-2] + table[i-1]
	}

	return table[n]
}

/*
Complexidade:
Espaco: O(1)
Tempo: O(n)
*/
func fibButtonUp2(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	curr, prev := 0, 1

	for i := 2; i < n+1; i++ {
		prev, curr = curr, prev+curr
	}

	return curr

}
