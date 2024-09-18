package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

/*
complixidade
Tempo: O(2**n)
Spaco: O(n)
*/
func generateParenthesis(n int) []string {

	response, solution := []string{}, []string{}

	var backtrack func(open, close int)

	backtrack = func(open, close int) {
		if len(solution) == 2*n {
			response = append(response, strings.Join(solution, ""))
			return
		}

		if open < n {
			solution = append(solution, "(")
			backtrack(open+1, close)
			solution = solution[:len(solution)-1]
		}

		if open > close {
			solution = append(solution, ")")
			backtrack(open, close+1)
			solution = solution[:len(solution)-1]
		}

	}

	backtrack(0, 0)
	return response
}
