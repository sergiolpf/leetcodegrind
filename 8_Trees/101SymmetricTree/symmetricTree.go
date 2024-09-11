package main

import "fmt"

func main() {
	//[4,2,7,1,3,6,9]
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(isSymmetric(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {

	return same(root, root)

}

func same(root1 *TreeNode, root2 *TreeNode) bool {

	if root1 == nil && root2 == nil {
		return true
	}

	if root1 != nil && root2 == nil ||
		root1 == nil && root2 != nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return same(root1.Left, root2.Right) && same(root1.Right, root2.Left)

}
