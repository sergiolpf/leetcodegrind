package main

import "fmt"

func main() {

	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}

/*
Tempo: O(n)
Espaco: O(1)
*/
func maxSubArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}

	currSum, maxSum := numbers[0], numbers[0]

	for i := 1; i < len(numbers); i++ {
		currSum = max(currSum+numbers[i], numbers[i])
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}
