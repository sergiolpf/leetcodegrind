package main

import "fmt"

func main() {
	fmt.Println(romanToInt("K"))
	//fmt.Println(romanToInt("LVIII"))
	//fmt.Println(romanToInt("MCMXCIV"))

}

func romanToInt(s string) int {
	var resultado int

	mapDeConversao := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1_000,
	}

	anterior := 0
	for ind := 0; ind < len(s); ind++ {
		valor := mapDeConversao[s[ind]]

		if valor > anterior {
			resultado += valor - 2*anterior
		} else {
			resultado += valor
		}

		anterior = valor

	}

	return resultado
}
