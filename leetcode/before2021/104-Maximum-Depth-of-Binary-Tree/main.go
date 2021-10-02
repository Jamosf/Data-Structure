package main

import (
	"math"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, depht *int) {
	if root == nil {
		*depht = 0
		return
	}
	var left, right int
	if root.Left != nil {
		dfs(root.Left, &left)
	}
	if root.Right != nil {
		dfs(root.Right, &right)
	}
	*depht = 1 + int(math.Max(float64(left), float64(right)))
}

func maxDepth(root *TreeNode) int {
	var r int
	dfs(root, &r)
	return r
}
