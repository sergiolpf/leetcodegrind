package main

import "fmt"

func main() {
	//[1,2,2,3,3,null,null,4,4]
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(diameterOfBinaryTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {

	if root.Left == nil && root.Right == nil {
		return 0
	}

	_, diameter := getMaxDiameter(root)
	return diameter
}

func getMaxDiameter(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	dLeft, diaLeft := getMaxDiameter(root.Left)
	dRight, diaRight := getMaxDiameter(root.Right)

	return max(dLeft, dRight) + 1, max((dLeft + dRight), diaLeft, diaRight)
}
