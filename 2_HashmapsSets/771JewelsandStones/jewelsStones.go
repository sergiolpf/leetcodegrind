package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
}

func numJewelsInStones(jewels string, stones string) int {

	var quantidade int

	jewelMap := make(map[rune]int)

	for _, value := range jewels {
		jewelMap[value] = 1
	}

	for _, stone := range stones {
		if _, ok := jewelMap[stone]; ok {
			quantidade++
		}
	}

	return quantidade
}
