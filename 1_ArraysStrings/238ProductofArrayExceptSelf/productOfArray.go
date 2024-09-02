package main

import "fmt"

/*
	 [1, 2, 3, 4]
pre- [1, 1, 2, 6]
suf- [24,12,4, 1]
res- [24,12,8, 6]

[0, 2]
[1, 0]
[2 ,1]
[2, 0]
*/

func main() {
	fmt.Println(productExceptSelf([]int{0, 2}))
}

func productExceptSelf(nums []int) []int {
	pre, suf := make([]int, len(nums)), make([]int, len(nums))
	pre[0] = 1
	suf[len(nums)-1] = 1

	if len(nums) == 2 {
		return []int{nums[1], nums[0]}
	}

	for ind := 1; ind < len(nums); ind++ {
		pre[ind] = pre[ind-1] * nums[ind-1]
	}
	fmt.Printf("%#v\n", pre)

	for ind := len(nums) - 2; ind >= 0; ind-- {
		suf[ind] = suf[ind+1] * nums[ind+1]
	}
	fmt.Printf("%#v\n", suf)

	resultado := make([]int, len(nums))
	for ind := 0; ind < len(nums); ind++ {
		resultado[ind] = pre[ind] * suf[ind]
	}

	return resultado
}
