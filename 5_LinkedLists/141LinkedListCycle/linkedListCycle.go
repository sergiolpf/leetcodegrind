package main

import (
	"fmt"
)

func main() {

	element := ListNode{
		Val:  2,
		Next: nil,
	}
	head := ListNode{
		Val:  3,
		Next: &element,
	}

	minhaLista := ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  4,
			Next: &element,
		},
	}

	element.Next = &minhaLista
	// minhaLista2 := ListNode{
	// 	Val:  2,
	// 	Next: &head,
	// }

	// head.Next = &minhaLista
	//head.Next = &minhaLista2

	fmt.Printf("%#v", hasCycle(&head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	dummy := &ListNode{}
	dummy.Next = head
	slow, fast := dummy, dummy

	if head == nil || head.Next == nil {
		return false
	}

	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
