package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	mapComplement := make(map[int]int, len(nums))

	for ind, value := range nums {
		mapComplement[value] = ind
	}

	for ind, value := range nums {
		comp := target - value
		if ind2, ok := mapComplement[comp]; ok {
			if ind2 != ind {
				return []int{ind, ind2}

			}
		}
	}

	return nil
}
