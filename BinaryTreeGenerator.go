package GolangUtilities

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateBinaryTree(val, maxDepth int) *TreeNode {
	if maxDepth == 0 {
		return nil
	}

	node := &TreeNode{Val: val}
	maxChildVal := int(math.Pow(2, float64(maxDepth-1)))

	if val*2 <= maxChildVal {
		node.Left = GenerateBinaryTree(val*2, maxDepth-1)
	}
	if val*2+1 <= maxChildVal {
		node.Right = GenerateBinaryTree(val*2+1, maxDepth-1)
	}

	return node
}
