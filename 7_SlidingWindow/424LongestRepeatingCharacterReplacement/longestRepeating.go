package main

import "fmt"

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
}

func characterReplacement(s string, k int) int {

	if len(s) <= 1 {
		return len(s)
	}

	left, right := 0, 0
	longest := 0
	mRunes := make(map[rune]int)
	sRunes := []rune(s)

	for right < len(s) {
		mRunes[sRunes[right]]++
		biggest := getMax(mRunes)

		for (right-left+1)-biggest > k {
			mRunes[sRunes[left]]--
			biggest = getMax(mRunes)
			left++

		}

		longest = max(longest, (right - left + 1))
		right++
	}

	return longest

}

func getMax(mapa map[rune]int) int {
	maxValue := 0
	for _, v := range mapa {
		maxValue = max(maxValue, v)
	}

	return maxValue
}
