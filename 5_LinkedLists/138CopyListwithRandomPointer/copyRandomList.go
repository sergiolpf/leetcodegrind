package main

import (
	"fmt"
)

func main() {

	minhaLista := Node{
		Val: 1,
		Next: &Node{
			Val: 1,
			Next: &Node{
				Val: 2,
				Next: &Node{
					Val: 3,
					Next: &Node{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Printf("%#v", copyRandomList(&minhaLista))
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	mapOfExistingToNewNodes := make(map[*Node]*Node)

	if head == nil {
		return nil
	}

	for aux := head; aux != nil; aux = aux.Next {

		if _, ok := mapOfExistingToNewNodes[aux]; !ok {
			mapOfExistingToNewNodes[aux] = &Node{Val: aux.Val}
		}
	}

	for aux := head; aux != nil; aux = aux.Next {
		mapOfExistingToNewNodes[aux].Next = mapOfExistingToNewNodes[aux.Next]
		mapOfExistingToNewNodes[aux].Random = mapOfExistingToNewNodes[aux.Random]
	}

	return mapOfExistingToNewNodes[head]

}
