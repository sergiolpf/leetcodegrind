package main

import (
	"container/heap"
	"fmt"
)

func main() {

	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

type IntegerHeap []int

func (h IntegerHeap) Len() int           { return len(h) }
func (h IntegerHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntegerHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntegerHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntegerHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	maxHeap := &IntegerHeap{}
	heap.Init(maxHeap)

	for _, v := range nums {
		heap.Push(maxHeap, v)
	}

	var res int

	for aux := 0; aux < k; aux++ {
		res = heap.Pop(maxHeap).(int)
	}

	return res
}
