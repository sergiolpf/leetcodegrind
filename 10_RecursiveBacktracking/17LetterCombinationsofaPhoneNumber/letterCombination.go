package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(letterCombinations("2"))
}

/*
 */
func letterCombinations(digits string) []string {

	if digits == "" {
		return nil
	}

	response, solution := []string{}, []string{}
	sDigits := []rune(digits)

	digitToLetter := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var backtrack func(ind int)
	backtrack = func(ind int) {

		if len(solution) == len(digits) {
			response = append(response, strings.Join(solution, ""))
			return
		}

		for _, v := range digitToLetter[sDigits[ind]] {

			solution = append(solution, v)
			backtrack(ind + 1)
			solution = solution[:len(solution)-1]

		}

	}

	backtrack(0)
	return response

}
