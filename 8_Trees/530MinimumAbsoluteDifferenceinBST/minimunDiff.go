package main

import (
	"fmt"
	"math"
)

func main() {
	//[4,2,7,1,3,6,9]
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			// Left: &TreeNode{
			// 	Val: 1,
			// },
			// Right: &TreeNode{
			// 	Val: 3,
			// },
		},
		Right: &TreeNode{
			Val: 48,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 49,
			},
		},
	}

	fmt.Println(getMinimumDifference(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var MinDiff, Previous int

/*
Performace
Time: O(n)
Space: O(n)
*/
func getMinimumDifference(root *TreeNode) int {
	Previous = math.MaxInt64
	MinDiff = math.MaxInt64

	inOrder(root)

	return MinDiff
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	inOrder(root.Left)

	if root.Val > Previous {
		MinDiff = min(MinDiff, root.Val-Previous)
	} else {
		MinDiff = min(MinDiff, Previous-root.Val)
	}

	Previous = root.Val

	inOrder(root.Right)

}
