package main

import "fmt"

func main() {
	//[1,2,3,6,9,8,7,4,5]
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
}

/*
[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
*/

func spiralOrder(matrix [][]int) []int {

	resultado := []int{}

	if len(matrix) == 1 && len(matrix[0]) == 1 {
		return append(resultado, matrix[0][0])
	}

	topo, direita, fundo, esquerda := 0, len(matrix[0])-1, len(matrix)-1, 0

	for {
		//direita
		for ind := esquerda; ind <= direita; ind++ {
			resultado = append(resultado, matrix[topo][ind])
		}
		topo++
		if topo > fundo {
			break
		}

		//baixo
		for ind := topo; ind <= fundo; ind++ {
			resultado = append(resultado, matrix[ind][direita])
		}
		direita--
		if esquerda > direita {
			break
		}

		//esquerda
		for ind := direita; ind >= esquerda; ind-- {
			resultado = append(resultado, matrix[fundo][ind])
		}
		fundo--
		if topo > fundo {
			break
		}

		//cima
		for ind := fundo; ind >= topo; ind-- {
			resultado = append(resultado, matrix[ind][esquerda])
		}
		esquerda++
		if esquerda > direita {
			break
		}

	}

	return resultado
}
