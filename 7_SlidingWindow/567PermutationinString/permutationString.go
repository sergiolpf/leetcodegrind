package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println(checkInclusion("hello", "ooolleoooleh"))
}

func checkInclusion(s1 string, s2 string) bool {

	if len(s2) < len(s1) {
		return false
	}
	mS1 := make(map[rune]int)
	left, right := 0, 0

	for _, v := range s1 {
		mS1[v]++
	}

	ss2 := []rune(s2)
	mapS2 := make(map[rune]int)
	for right < len(s2) {
		for ; right-left+1 <= len(s1); right++ {
			mapS2[ss2[right]]++
		}
		if reflect.DeepEqual(mS1, mapS2) {
			return true
		}
		if mapS2[ss2[left]] == 1 {
			delete(mapS2, ss2[left])
		} else {
			mapS2[ss2[left]]--
		}
		left++
	}

	return false
}
