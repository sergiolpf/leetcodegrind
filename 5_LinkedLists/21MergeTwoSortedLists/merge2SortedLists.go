package main

import "fmt"

func main() {

	minhaLista := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	minhaLista2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	fmt.Printf("%#v", mergeTwoLists(&minhaLista, &minhaLista2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var mergedList *ListNode
	var mcur *ListNode

	if list1 == nil && list2 == nil {
		return mergedList
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var cur1, cur2 *ListNode
	for cur1, cur2 = list1, list2; cur1 != nil && cur2 != nil; {

		if cur1.Val <= cur2.Val {
			if mergedList == nil {
				mergedList = cur1
				mcur = mergedList
			} else {
				mcur.Next = cur1
				mcur = cur1
			}
			cur1 = cur1.Next
		} else {
			if mergedList == nil {
				mergedList = cur2
				mcur = mergedList
			} else {
				mcur.Next = cur2
				mcur = cur2
			}
			cur2 = cur2.Next
		}
	}

	if cur1 != nil {
		mcur.Next = cur1
	}

	if cur2 != nil {
		mcur.Next = cur2
	}

	return mergedList
}
