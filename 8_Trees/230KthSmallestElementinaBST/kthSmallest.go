package main

import "fmt"

func main() {
	//[4,2,7,1,3,6,9]
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	fmt.Println(kthSmallest(root, 6))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Complexity:
Time: O(n)
Space: O(n)
*/
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	elementsInOrder := getInOrderTraverse(root)

	return elementsInOrder[k-1]
}

func getInOrderTraverse(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	response := []int{}
	left := getInOrderTraverse(root.Left)
	right := getInOrderTraverse(root.Right)

	response = append(response, left...)
	response = append(response, root.Val)
	response = append(response, right...)

	return response
}
