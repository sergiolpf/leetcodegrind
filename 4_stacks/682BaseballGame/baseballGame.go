package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}

func calPoints(operations []string) int {
	calculation := []int{}

	for _, op := range operations {
		switch op {
		case "+":
			calculation = append(calculation, calculation[len(calculation)-1]+calculation[len(calculation)-2])
		case "D":

			calculation = append(calculation, calculation[len(calculation)-1]*2)
		case "C":
			calculation = calculation[:len(calculation)-1]
		default:
			intScore, err := strconv.Atoi(op)
			if err != nil {
				// Handle the error, for example:
				fmt.Println("Error converting string to int:", err)
				return 0
			}
			calculation = append(calculation, intScore)
		}
	}

	sum := 0

	for _, value := range calculation {
		sum += value
	}

	return sum
}
