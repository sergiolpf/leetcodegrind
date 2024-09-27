package main

import "fmt"

func main() {
	fmt.Println(uniquePaths3(3, 7))
}

/*
Tempo: O(2^ (m*n))
Espaco: O(n)
*/
func uniquePaths(m int, n int) int {

	if m == 1 || n == 1 {
		return 1
	}

	var helper func(x, y int) int
	helper = func(x, y int) int {

		if x == 1 || y == 1 {
			return 1
		}

		return helper(x, y-1) + helper(x-1, y)
	}

	return helper(m, n)

}

/*
Tempo: O(n)
Espaco: O(n)
*/
func uniquePaths2(m int, n int) int {

	memo := make(map[[2]int]int)
	memo[[2]int{0, 0}] = 1

	var helper func(x, y int) int
	helper = func(x, y int) int {

		if memo[[2]int{x, y}] > 0 {
			return memo[[2]int{x, y}]
		}
		if x < 0 || y < 0 || x == m || y == n {
			return 0
		}

		val := helper(x, y-1) + helper(x-1, y)
		memo[[2]int{x, y}] = val
		return val
	}

	return helper(m-1, n-1)

}

/*
Tempo: O(n)
Espaco: O(n)
*/
func uniquePaths3(m int, n int) int {

	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}

	memo[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}

			val := 0
			if i > 0 {
				val += memo[i-1][j]
			}
			if j > 0 {
				val += memo[i][j-1]
			}

			memo[i][j] = val
		}

	}

	return memo[m-1][n-1]
}
