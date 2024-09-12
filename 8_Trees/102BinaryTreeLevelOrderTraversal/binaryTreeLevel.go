package main

import (
	"fmt"
)

func main() {
	//[4,2,7,1,3,6,9]
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: -1,
				Right: &TreeNode{
					Val: -8,
				},
			},
		},
	}

	fmt.Println(levelOrder(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
Time: O(n)
Space: O(n)
*/
func levelOrder(root *TreeNode) [][]int {

	queue := [][]*TreeNode{}
	response := [][]int{}

	if root == nil {
		return nil
	}

	queue = append(queue, []*TreeNode{root})
	response = append(response, []int{root.Val})

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		resValue := []int{}
		new := []*TreeNode{}
		for _, v := range top {

			if v.Left != nil {
				new = append(new, v.Left)
				resValue = append(resValue, v.Left.Val)
			}
			if v.Right != nil {
				new = append(new, v.Right)
				resValue = append(resValue, v.Right.Val)
			}

		}
		if len(new) > 0 {
			queue = append(queue, new)
			response = append(response, resValue)
		}

	}

	return response
}
