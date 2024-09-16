package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(kClosest2([][]int{
		{3, 3},
		{5, -1},
		{-2, 4},
	}, 2))
}

type Item struct {
	ponto []int
	dist  float64
}

type coordenadasStack []*Item

func (h coordenadasStack) Len() int           { return len(h) }
func (h coordenadasStack) Less(i, j int) bool { return h[i].dist > h[j].dist }
func (h coordenadasStack) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *coordenadasStack) Push(x any) {
	item := &Item{
		ponto: x.([]int),
		dist:  math.Pow(float64(x.([]int)[0]), 2) + math.Pow(float64(x.([]int)[1]), 2),
	}
	*h = append(*h, item)
}

func (h *coordenadasStack) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x.ponto
}

/*
Complexidade
Tempo: O(NLogn) - a parte do sort/push/pop eh logN e como fazemos isso n vezes, logo NLogn
Espaco: O(n)
Para que a funcao abaixo funcione, temos que fazer a funcao less sendo rodenada do menor para o maior
---Cheque que a funcao Less usada seja a da proxima linha.
func (h coordenadasStack) Less(i, j int) bool { return h[i].dist < h[j].dist }
*/
func kClosest(points [][]int, k int) [][]int {

	sortedByDistance := &coordenadasStack{}

	heap.Init(sortedByDistance)

	for _, v := range points {
		heap.Push(sortedByDistance, v)
	}

	response := [][]int{}

	for i := 0; i < k; i++ {
		res := heap.Pop(sortedByDistance).([]int)
		response = append(response, res)

	}

	return response
}

/*
Complexidade
Tempo: O(NLogK) - a parte do sort/push/pop eh logK - a estrutura sempre tera somente K elementos. e como fazemos isso N vezes, logo NLogK
Espaco: O(n)
Para que a funcao abaixo funcione, temos que fazer a funcao less sendo rodenada do menor para o maior
---Cheque que a funcao Less usada seja a da proxima linha.
func (h coordenadasStack) Less(i, j int) bool { return h[i].dist < h[j].dist }
*/
func kClosest2(points [][]int, k int) [][]int {
	sortedByDistance := &coordenadasStack{}

	heap.Init(sortedByDistance)

	for ind, v := range points {

		heap.Push(sortedByDistance, v)
		if ind >= k {
			heap.Pop(sortedByDistance)
		}
	}

	response := [][]int{}

	for i := 0; i < k; i++ {
		res := heap.Pop(sortedByDistance).([]int)
		response = append(response, res)

	}

	return response

}
