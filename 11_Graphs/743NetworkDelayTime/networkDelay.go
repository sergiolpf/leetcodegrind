package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(networkDelayTime([][]int{
		{1, 2, 1},
		{2, 3, 2},
		{1, 3, 4},
	}, 3, 1))
}

type Item struct {
	time int
	dest int
}

type minHeap []Item

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *minHeap) Pop() any {
	old := *h
	size := len(old)
	x := old[size-1]
	*h = old[:size-1]
	return x

}
func networkDelayTime(times [][]int, n int, k int) int {

	graph := make(map[int][]Item)

	for _, tupla := range times {
		graph[tupla[0]] =
			append(graph[tupla[0]],
				Item{time: tupla[2], dest: tupla[1]})

	}

	localMinHeap := minHeap{}
	heap.Init(&localMinHeap)
	heap.Push(&localMinHeap, Item{
		dest: k, time: 0,
	})
	minTime := make(map[int]int)

	for len(localMinHeap) > 0 {
		item := heap.Pop(&localMinHeap).(Item)

		if _, ok := minTime[item.dest]; ok {
			continue
		}

		minTime[item.dest] = item.time

		for _, neighbors := range graph[item.dest] {
			if _, ok := minTime[neighbors.dest]; !ok {
				heap.Push(&localMinHeap, Item{dest: neighbors.dest, time: item.time + neighbors.time})
			}
		}
	}

	if len(minTime) == n {
		maxValue := math.MinInt64
		for _, v := range minTime {
			maxValue = max(maxValue, v)
		}
		return maxValue
	}

	return -1
}
