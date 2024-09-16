package main

import (
	"fmt"
	"math"
)

func main() {
	//[1,2,2,3,3,null,null,4,4]
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			// Left: &TreeNode{
			// 	Val: 3,
			// 	Left: &TreeNode{
			// 		Val: 4,
			// 	},
			// 	Right: &TreeNode{
			// 		Val: 4,
			// 	},
			// },
			// Right: &TreeNode{
			// 	Val: 3,
			// },
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(isValidBST(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
complexidade
Tempo: O(n)
Espaco: O(h) - h < n - altura da arvore
*/
func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *TreeNode, minn, maxx int) bool {

	if root == nil {
		return true
	}

	if root.Val <= minn || root.Val >= maxx {
		return false
	}

	return isValid(root.Left, minn, root.Val) && isValid(root.Right, root.Val, maxx)
}
