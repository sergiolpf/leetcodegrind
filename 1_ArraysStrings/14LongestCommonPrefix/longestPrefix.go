package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	maiorPrefixo := strs[0]
	isprefixo := true

	for len(maiorPrefixo) > 0 {
		for ind := 1; ind < len(strs); ind++ {
			if !strings.HasPrefix(strs[ind], maiorPrefixo) {
				isprefixo = false
				break
			}
		}
		if isprefixo {
			break
		}

		maiorPrefixo = maiorPrefixo[:len(maiorPrefixo)-1]
		isprefixo = true

	}
	return maiorPrefixo

}
