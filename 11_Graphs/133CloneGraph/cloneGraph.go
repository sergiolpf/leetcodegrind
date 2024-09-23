package main

func main() {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

/*
Complexidade:
Tempo: O(V+E)
Espaco: O(v)
*/
func cloneGraph(node *Node) *Node {

	if node == nil {
		return nil
	}

	clonedMap := make(map[*Node]*Node)
	stack := []*Node{}

	clonedMap[node] = &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}
	stack = append(stack, node)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, next := range curr.Neighbors {
			if _, ok := clonedMap[next]; !ok {
				stack = append(stack, next)
				clonedMap[next] = &Node{
					Val:       next.Val,
					Neighbors: []*Node{},
				}
			}
		}
	}

	for original, clone := range clonedMap {
		for _, next := range original.Neighbors {
			clone.Neighbors = append(clone.Neighbors, clonedMap[next])
		}

	}

	return clonedMap[node]

}
