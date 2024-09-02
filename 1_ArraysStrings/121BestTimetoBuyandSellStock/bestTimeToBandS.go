package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {

	if len(prices) == 1 {
		return 0
	}

	valorCompra := prices[0]
	lucro := 0

	for _, preco := range prices {
		if preco < valorCompra {
			valorCompra = preco
			continue
		}

		if preco-valorCompra > lucro {
			lucro = preco - valorCompra
		}

	}

	return lucro
}
