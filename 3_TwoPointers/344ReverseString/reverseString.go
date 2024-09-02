package main

import "fmt"

func main() {
	str := []byte{
		'H', 'a', 'n', 'n', 'a', 'h',
	}
	reverseString(str)
	fmt.Println(string(str))
}

func reverseString(s []byte) {

	if len(s) == 1 {
		return
	}

	for left, right := 0, len(s)-1; left <= right; right, left = right-1, left+1 {
		s[left], s[right] = s[right], s[left]
	}

}
