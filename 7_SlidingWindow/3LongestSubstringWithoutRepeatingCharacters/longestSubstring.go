package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("au"))
}

func lengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}
	left, right := 0, 0
	mapUnique := make(map[rune]bool, len(s))
	sRune := []rune(s)
	longest := 0

	for right < len(sRune) {
		if !mapUnique[sRune[right]] {
			mapUnique[sRune[right]] = true
			right++
		} else {
			for mapUnique[sRune[right]] {
				delete(mapUnique, sRune[left])
				left++
			}

		}
		longest = max(longest, len(mapUnique))
	}

	return longest

}
