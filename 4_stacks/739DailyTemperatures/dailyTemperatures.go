package main

import (
	"fmt"
)

func main() {
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}

type element struct {
	temp int
	ind  int
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	littleStack := []element{}

	if len(temperatures) == 1 {
		return []int{0}
	}

	littleStack = append(littleStack, element{temp: temperatures[0], ind: 0})

	for ind := 1; ind < len(temperatures); {
		value := temperatures[ind]

		if len(littleStack) > 0 {

			pop := littleStack[len(littleStack)-1]

			if pop.temp >= value {

				littleStack = append(littleStack, element{temp: value, ind: ind})
				ind++

			} else {
				result[pop.ind] = ind - pop.ind
				littleStack = littleStack[:len(littleStack)-1]
			}

		} else {
			littleStack = append(littleStack, element{temp: value, ind: ind})
			ind++
		}
	}

	return result
}
