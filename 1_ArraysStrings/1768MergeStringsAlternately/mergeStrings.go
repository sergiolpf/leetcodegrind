package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))

}

func mergeAlternately(word1 string, word2 string) string {

	var resultado []string

	rWord1 := []rune(word1)
	rWord2 := []rune(word2)

	indw1, indw2 := 0, 0

	for indw1 < len(word1) && indw2 < len(word2) {
		resultado = append(resultado, string(rWord1[indw1]))
		resultado = append(resultado, string(rWord2[indw2]))
		indw1++
		indw2++
	}

	if indw1 < len(word1) {
		resultado = append(resultado, string(rWord1[indw1:]))
	}

	if indw2 < len(word2) {
		resultado = append(resultado, string(rWord2[indw2:]))
	}

	return strings.Join(resultado, "")
}
