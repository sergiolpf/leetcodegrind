package main

func main() {

}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	if n == 1 {
		return 1
	}

	left := 0
	right := n - 1
	pivot := 0

	for left <= right {
		pivot = left + (right-left)/2

		if isBadVersion(pivot) {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return left

}
