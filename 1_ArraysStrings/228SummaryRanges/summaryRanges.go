package main

import "fmt"

func main() {
	fmt.Println(summaryRanges([]int{-1}))
}

func summaryRanges(nums []int) []string {

	var resultado []string

	if len(nums) == 0 {
		return resultado
	}
	if len(nums) == 1 {
		return append(resultado, fmt.Sprintf("%v", nums[0]))
	}

	mapaDeRanges := make(map[int]int)
	valorAtual := nums[0]
	mapaDeRanges[valorAtual] = valorAtual

	for ind := 1; ind < len(nums); ind++ {
		if nums[ind]-mapaDeRanges[valorAtual] == 1 {
			mapaDeRanges[valorAtual] = nums[ind]
		} else {
			valorAtual = nums[ind]
			mapaDeRanges[valorAtual] = nums[ind]
		}
	}

	for _, v := range nums {
		valor, ok := mapaDeRanges[v]

		if ok {
			if valor == v {
				resultado = append(resultado, fmt.Sprintf("%v", v))
			} else {
				resultado = append(resultado, fmt.Sprintf("%v->%v", v, valor))
			}
		}
	}

	return resultado
}
