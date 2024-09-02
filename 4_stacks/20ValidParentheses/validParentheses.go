package main

import "fmt"

func main() {
	fmt.Println(isValid("){"))
}

func isValid(s string) bool {

	stack := []rune{}

	if len(s) == 1 {
		return false
	}

	for _, value := range s {
		switch value {
		case '(', '[', '{':
			stack = append(stack, value)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		}
	}

	return len(stack) == 0
}
