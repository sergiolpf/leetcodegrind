package main

import "fmt"

func main() {

	fmt.Println(isSubsequence("b", "abc"))

}

func isSubsequence(s, t string) bool {

	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	lIndex, rIndex := 0, 0

	count := 0
	for rIndex < len(t) && lIndex < len(s) {
		if s[lIndex] == t[rIndex] {
			count++
			lIndex++
		}

		rIndex++
	}

	return count == len(s)
}
