package main

import "fmt"

func main() {
	//[1,2,2,3,3,null,null,4,4]
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 0,
				},
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	subRoot := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(isSubtree(root, subRoot))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if subRoot == nil {
		return true
	}

	if root.Val == subRoot.Val {
		if isSameTree(root, subRoot) {
			return true
		}
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)

}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p != nil && q == nil) ||
		p == nil && q != nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}
