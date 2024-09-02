package main

import "fmt"

func main() {

	minhaLista := ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	fmt.Printf("%#v", removeNthFromEnd(&minhaLista, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head.Next == nil && n == 1 {
		return nil
	}

	count := 0
	for aux := head; aux != nil; aux = aux.Next {
		count++
	}

	count -= n

	if count == 0 {
		if head.Next == nil {
			return nil
		} else {
			return head.Next
		}
	}

	previous := head
	for aux, contador := head, 0; aux != nil; contador++ {

		if contador == count {
			if aux.Next != nil {
				previous.Next = aux.Next
			} else {
				previous.Next = nil
			}
			return head
		}
		previous = aux
		aux = aux.Next
	}

	return head
}
