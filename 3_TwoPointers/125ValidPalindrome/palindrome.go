package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {

	filtered := ""

	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			filtered += strings.ToLower(string(v))
		}
	}

	for left, right := 0, len(filtered)-1; left <= right; left, right = left+1, right-1 {
		if filtered[left] != filtered[right] {
			return false
		}
	}

	return true
}
