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

	fmt.Println(lowestCommonAncestor(root, &TreeNode{
		Val: 7}, &TreeNode{
		Val: 1}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Complexidade:
Tempo: O(h) - h < n - altura da arvore
Espaco: O(h) - h < n - altura da arvore
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	if (root.Val < p.Val && root.Val > q.Val) || (root.Val > p.Val && root.Right.Val < q.Val) {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
