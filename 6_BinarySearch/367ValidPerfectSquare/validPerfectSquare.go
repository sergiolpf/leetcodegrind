package main

import "fmt"

func main() {

	fmt.Println(isPerfectSquare(2))
}

func isPerfectSquare(num int) bool {

	if num == 1 {
		return true
	}

	left := 0
	right := num
	pivot := 0
	sqr := 0

	for left <= right {
		pivot = left + (right-left)/2
		sqr = pivot * pivot

		if sqr == num {
			return true
		}

		if sqr < num {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return false
}
