package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(tokens []string) int {

	values := []int{}

	for _, op := range tokens {
		switch op {
		case "+":
			if len(values) > 1 {
				op1 := values[len(values)-1]
				op2 := values[len(values)-2]
				values = values[:len(values)-2]
				values = append(values, op2+op1)
			}
		case "-":
			if len(values) > 1 {
				op1 := values[len(values)-1]
				op2 := values[len(values)-2]
				values = values[:len(values)-2]
				values = append(values, op2-op1)
			}
		case "*":
			if len(values) > 1 {
				op1 := values[len(values)-1]
				op2 := values[len(values)-2]
				values = values[:len(values)-2]
				values = append(values, op2*op1)
			}
		case "/":
			if len(values) > 1 {
				op1 := values[len(values)-1]
				op2 := values[len(values)-2]
				values = values[:len(values)-2]
				values = append(values, op2/op1)
			}
		default:
			intScore, err := strconv.Atoi(op)
			if err != nil {
				// Handle the error, for example:
				fmt.Println("Error converting string to int:", err)
				return 0
			}
			values = append(values, intScore)

		}

	}
	return values[0]
}
