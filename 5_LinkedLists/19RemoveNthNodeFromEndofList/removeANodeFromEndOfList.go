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

	dummy := &ListNode{}
	dummy.Next = head
	ahead, behind := dummy, dummy

	for aux := 1; aux <= n+1; aux++ {
		ahead = ahead.Next
	}

	for ; ahead != nil; ahead, behind = ahead.Next, behind.Next {

	}
	behind.Next = behind.Next.Next

	return dummy.Next
}
