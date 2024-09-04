package main

import "fmt"

func main() {

	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 70},
	}

	fmt.Println(searchMatrix(matrix, 11))
}

func searchMatrix(matrix [][]int, target int) bool {
	up := 0
	down := len(matrix) - 1
	pivot := 0

	if len(matrix) == 1 && len(matrix[0]) == 1 {
		if matrix[0][0] == target {
			return true
		} else {
			return false
		}

	}

	for up <= down {
		pivot = up + (down-up)/2

		if matrix[pivot][0] == target {
			return true
		}

		if matrix[pivot][0] > target {
			down = pivot - 1
		} else {
			up = pivot + 1
		}
	}
	if down < 0 {
		return false
	}

	left := 0
	right := len(matrix[0]) - 1

	for left <= right {
		pivot2 := left + (right-left)/2

		if matrix[down][pivot2] == target {
			return true
		}
		if matrix[down][pivot2] > target {
			right = pivot2 - 1
		} else {
			left = pivot2 + 1
		}

	}

	return false

}
