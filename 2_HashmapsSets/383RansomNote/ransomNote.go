package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {

	magazineLetters := make(map[rune]int)

	for _, note := range magazine {
		magazineLetters[note]++
	}

	for _, ransomLetter := range ransomNote {
		magazineLetters[ransomLetter]--
		if magazineLetters[ransomLetter] < 0 {
			return false
		}
	}

	return true
}
