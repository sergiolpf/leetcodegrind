package main

func main() {

}

/*
Complexidade
Tempo: O(2^n)
Espaco: O(n)
*/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	return climbStairs(n-2) + climbStairs(n-1)
}

/*
Complexidade
Tempo: O(2^n)
Espaco: O(n)
*/
func climbStairsMemoRec(n int) int {
	if n <= 2 {
		return n
	}

	memo := make(map[int]int)
	memo[1] = 1
	memo[2] = 2

	var dp func(x int) int
	dp = func(x int) int {

		if memo[x] > 0 {
			return memo[x]
		}

		memo[x] = dp(x-2) + dp(x-1)
		return memo[x]
	}

	return dp(n)
}

/*
Complexidade
Tempo: O(2^n)
Espaco: O(n)
*/
func climbStairsMemoNonRec(n int) int {
	if n <= 2 {
		return n
	}

	memo := make([]int, n+1)
	memo[1] = 1
	memo[2] = 2

	for x := 3; x < n+1; x++ {
		memo[x] = memo[x-1] + memo[x-2]
	}

	return memo[n]
}

/*
Complexidade
Tempo: O(2^n)
Espaco: O(1)
*/
func climbStairsMemoNonRecConstantSpace(n int) int {
	if n <= 2 {
		return n
	}

	previous, current := 1, 2

	for x := 3; x < n+1; x++ {
		previous, current = current, current+previous
	}

	return current
}
