package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

type Item struct {
	value int
	qtd   int
}

type IntPriorityHeap []*Item

func (h IntPriorityHeap) Len() int { return len(h) }
func (h IntPriorityHeap) Less(i, j int) bool {
	return h[i].qtd > h[j].qtd
}
func (h IntPriorityHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntPriorityHeap) Push(x any) {
	*h = append(*h, x.(*Item))
}

func (h *IntPriorityHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
Complexidade
Tempo: O(nLogK) onde N eh o tamanho do array nums e K eh o element kth
*/
func topKFrequent(nums []int, k int) []int {
	maxList := &IntPriorityHeap{}
	maplist := make(map[int]int)

	heap.Init(maxList)

	for _, v := range nums {
		maplist[v]++

	}
	for ind, v := range maplist {
		item := &Item{
			value: ind,
			qtd:   v,
		}
		heap.Push(maxList, item)
	}

	response := []int{}

	for aux := 0; aux < k; aux++ {
		item := heap.Pop(maxList).(*Item)

		response = append(response, item.value)
	}
	return response
}
