package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 3, 3}, 6))
}

/*
Example 1:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
Example 2:

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
Example 3:

Input: numbers = [-1,0], target = -1
Output: [1,2]
*/
func twoSum(numbers []int, target int) []int {
	resultado := []int{}

	for left, right := 0, len(numbers)-1; left < right; {
		if numbers[left]+numbers[right] == target {
			resultado = append(resultado, left+1, right+1)
			break
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}

	}
	return resultado

}
