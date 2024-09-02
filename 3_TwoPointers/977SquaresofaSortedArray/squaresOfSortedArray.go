package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortedSquares1([]int{-7, -3, 2, 3, 11}))
}

func sortedSquares(nums []int) []int {
	resultado := []int{}

	for _, value := range nums {
		resultado = append(resultado, value*value)
	}

	sort.Ints(resultado)
	return resultado
}

func sortedSquares1(nums []int) []int {

	if len(nums) == 1 {
		return []int{nums[0] * nums[0]}
	}

	resultado := []int{}
	unsortedResults := []int{}

	for _, value := range nums {
		unsortedResults = append(unsortedResults, value*value)
	}

	L, R := 0, len(nums)-1

	for L <= R {
		if unsortedResults[L] >= unsortedResults[R] {
			resultado = append(resultado, unsortedResults[L])
			L++
		} else {
			resultado = append(resultado, unsortedResults[R])
			R--
		}
	}

	sorted := []int{}

	for ind := len(resultado) - 1; ind >= 0; ind-- {
		sorted = append(sorted, resultado[ind])
	}

	return sorted
}
