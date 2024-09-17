package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// list1 := ListNode{
	// 	Val: -2,
	// 	Next: &ListNode{
	// 		Val: -1,
	// 		Next: &ListNode{
	// 			Val: -1,
	// 			Next: &ListNode{
	// 				Val:  -1,
	// 				Next: nil,
	// 			},
	// 		},
	// 	},
	// }

	list2 := ListNode{
		Val:  1,
		Next: nil,
	}
	list3 := ListNode{
		Val:  0,
		Next: nil,
	}
	// list3 := ListNode{
	// 	Val: 2,
	// 	Next: &ListNode{
	// 		Val:  6,
	// 		Next: nil,
	// 	},
	// }

	minhaLista := []*ListNode{}

	minhaLista = append(minhaLista, &list2, &list3)
	minhaMergedList := mergeKLists(minhaLista)

	for minhaMergedList != nil {
		fmt.Printf("%v,", minhaMergedList.Val)
		minhaMergedList = minhaMergedList.Next
	}

	//fmt.Printf("%#v", mergeKLists(minhaLista))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type sortedHeap []*ListNode

func (h sortedHeap) Len() int { return len(h) }
func (h sortedHeap) Less(i, j int) bool {
	if h[i] == nil || h[j] == nil {
		if h[i] == nil {
			return false
		} else {
			return true
		}
	}
	return h[i].Val < h[j].Val
}
func (h sortedHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *sortedHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *sortedHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
Complexity
Tempo: O(NlogK) sendo K o tamanho da lista inicial
Espaco: O(N)
*/

func mergeKLists(lists []*ListNode) *ListNode {

	sequence := &sortedHeap{}
	heap.Init(sequence)

	for _, nodes := range lists {
		if nodes != nil {
			heap.Push(sequence, nodes)
		}
	}

	dummy := &ListNode{}
	curr := dummy

	for len(*sequence) > 0 {
		pop := heap.Pop(sequence).(*ListNode)
		curr.Next = pop
		curr = curr.Next
		if pop != nil && pop.Next != nil {
			heap.Push(sequence, pop.Next)
		}
	}

	return dummy.Next
}
