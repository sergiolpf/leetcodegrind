package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3}))
}

func containsDuplicate(nums []int) bool {

	uniqueList := make(map[int]int, len(nums))

	for _, num := range nums {
		uniqueList[num]++
		if uniqueList[num] == 2 {
			return true
		}

	}

	return false
}
