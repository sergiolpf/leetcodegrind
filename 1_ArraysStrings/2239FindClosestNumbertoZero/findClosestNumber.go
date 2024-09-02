/*

Given an integer array nums of size n, return the number with the value closest to 0 in nums. If there are multiple answers, return the number with the largest value.

Example 1:

Input: nums = [-4,-2,1,4,8]
Output: 1
Explanation:
The distance from -4 to 0 is |-4| = 4.
The distance from -2 to 0 is |-2| = 2.
The distance from 1 to 0 is |1| = 1.
The distance from 4 to 0 is |4| = 4.
The distance from 8 to 0 is |8| = 8.
Thus, the closest number to 0 in the array is 1.
Example 2:

Input: nums = [2,-1,1]
Output: 1
Explanation: 1 and -1 are both the closest numbers to 0, so 1 being larger is returned.


Constraints:

1 <= n <= 1000
-105 <= nums[i] <= 105
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(closestNumber([]int{2, -1, 1}))
}

func closestNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	currNumber := nums[0]
	disTo0 := math.Abs(float64(nums[0] - 0))

	for _, value := range nums {
		if math.Abs(float64(value-0)) < disTo0 {
			currNumber = value
			disTo0 = math.Abs(float64(value - 0))
		} else if math.Abs(float64(value-0)) == disTo0 {
			if value > currNumber {
				currNumber = value
			}
		}

	}
	return currNumber
}
