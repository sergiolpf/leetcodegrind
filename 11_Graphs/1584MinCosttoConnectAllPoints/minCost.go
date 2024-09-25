package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCostConnectPoints([][]int{
		{0, 0},
		{2, 2},
		{3, 10},
		{5, 2},
		{7, 0},
	}))
}

type Item struct {
	point  int
	weight int
}

type sortedHeap []Item

func (h sortedHeap) Len() int { return len(h) }
func (h sortedHeap) Less(i, j int) bool {
	return h[i].weight < h[j].weight
}
func (h sortedHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *sortedHeap) Push(x any) {
	*h = append(*h, x.(Item))
}
func (h *sortedHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func manhattanDistance(p1, p2 []int) int {
	return int(math.Abs(float64(p1[0]-p2[0])) + math.Abs(float64(p1[1]-p2[1])))
}

/*
Complexidade
Tempo: O(n2 * logN) - o 2o for acaba fazendo isso ser n**2 - n sendo o tamanho do numero de pontos. - LogN por causa do priority heap, pra ordernar eh logn)
Espaco: O(n^2)

essa eh uma versao do algoritmo Prim que resolver a MST - Minimum Spanning Tree
*/
func minCostConnectPoints(points [][]int) int {

	if len(points) == 1 && len(points[0]) == 1 {
		return 0
	}

	minHeap := sortedHeap{}
	heap.Init(&minHeap)
	total_result := 0

	seen := make(map[int]bool, len(points)-1)
	heap.Push(&minHeap, Item{weight: 0, point: 0})

	for len(seen) < len(points) {
		localMan := heap.Pop(&minHeap).(Item)
		if seen[localMan.point] {
			continue
		}

		seen[localMan.point] = true

		total_result += localMan.weight

		for i := 0; i < len(points); i++ {
			if !seen[i] {
				weight := manhattanDistance(points[localMan.point], points[i])
				heap.Push(&minHeap, Item{point: i, weight: weight})
			}

		}

	}

	return total_result

}
