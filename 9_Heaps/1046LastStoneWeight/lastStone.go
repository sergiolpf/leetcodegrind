package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func main() {

	fmt.Println(lastStoneWeightHeap([]int{9, 10, 4, 5, 7, 1}))
}

/*
nesse metodo uso sort.Slice para ordenar decrescente toda vez q removo as 2 pedras
e adiciono o resto (se necessario)
complexidade:
Tempo: O(n2) - quadratico
Spaco: O(1) - nao uso nada novo, so organizo o array existente.
*/
func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	for len(stones) > 1 {

		sort.Slice(stones, func(i, j int) bool {
			return stones[i] > stones[j]
		})
		first := stones[0]
		sec := stones[1]
		res := int(math.Abs(float64(first - sec)))
		stones = append(stones, res)
		stones = stones[2:]
	}

	return stones[0]

}

type intPriorityHeap []int

func (h intPriorityHeap) Len() int           { return len(h) }
func (h intPriorityHeap) Less(i, j int) bool { return (h[i] > h[j]) }
func (h intPriorityHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intPriorityHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *intPriorityHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
A funcao abaixo usa o heap,
Complexidade:
Tempo: O(nlogn) - a parte do sort eh logn e como fazemos n vezes entao vira n.logn
Espaco: O(n)
*/
func lastStoneWeightHeap(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	maxHeap := &intPriorityHeap{}
	heap.Init(maxHeap)

	for _, v := range stones {
		heap.Push(maxHeap, v)
	}

	for maxHeap.Len() > 1 {
		first := heap.Pop(maxHeap).(int)
		second := heap.Pop(maxHeap).(int)

		if first != second {
			rest := int(math.Abs(float64(first - second)))
			heap.Push(maxHeap, rest)
		}
	}

	if len(*maxHeap) > 0 {
		return heap.Pop(maxHeap).(int)
	}
	return 0
}
