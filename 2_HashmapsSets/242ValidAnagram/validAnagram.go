package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("rat", "car"))
}

func isAnagram(s string, t string) bool {

	if len(s) == 1 || len(t) == 1 {
		return s == t
	}

	if len(s) != len(t) {
		return false
	}

	left := make(map[rune]int, len(s))
	right := make(map[rune]int, len(t))

	for _, letter := range s {
		left[letter]++
	}
	for _, letter := range t {
		right[letter]++
	}

	for letter, qtd := range left {
		if right[letter] != qtd {
			return false
		}
	}

	return true

}
