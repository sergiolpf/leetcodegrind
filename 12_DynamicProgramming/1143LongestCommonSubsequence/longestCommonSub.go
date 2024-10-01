package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("oxcpqrsvwf", "shmtulqrypy"))
}

/*
Tempo: O(m*n^2)
Espaco: O(M*N)
*/
func longestCommonSubsequence(text1 string, text2 string) int {

	m := len(text1)
	n := len(text2)
	memo := make([][]int, m+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	var memoSolve func(i, j int) int
	memoSolve = func(i, j int) int {
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		option1 := memoSolve(i+1, j)

		option2 := 0
		for x := j; x < n; x++ {
			if text1[i] == text2[x] {
				option2 = 1 + memoSolve(i+1, x+1)
				break
			}
		}

		memo[i][j] = max(option1, option2)
		return memo[i][j]
	}

	return memoSolve(0, 0)

}
