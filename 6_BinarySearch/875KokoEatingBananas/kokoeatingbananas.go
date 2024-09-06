package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
}

func minEatingSpeed(piles []int, h int) int {

	sort.Ints(piles)

	left := 1
	right := piles[len(piles)-1]

	for left < right {
		mid := (left + right) / 2

		if isOk(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}

	}

	return left

}

func isOk(piles []int, speed, hours int) bool {

	consumed := 0
	for _, v := range piles {
		coisa := int(math.Ceil(float64(v) / float64(speed)))
		consumed += coisa
	}

	return consumed <= hours
}
