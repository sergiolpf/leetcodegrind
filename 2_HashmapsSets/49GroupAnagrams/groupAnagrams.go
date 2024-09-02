package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	/*
		hash Set - key: strs[i] value: slice ordered of the key.
		then I compare all the
	*/
	orderedMap := make(map[string][]string, len(strs))

	for _, str := range strs {
		s_str := []rune(str)
		sort.Slice(s_str, func(i, j int) bool {
			return s_str[i] < s_str[j]
		})

		orderedMap[string(s_str)] = append(orderedMap[string(s_str)], str)
	}

	result := make([][]string, len(orderedMap))
	ind := 0
	for _, value := range orderedMap {
		result[ind] = value
		ind++
	}
	return result

}
