package main

import (
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkBalance(root *TreeNode, depth *int) bool {
	if root == nil {
		*depth = 0
		return true
	}
	var leftDepth, rightDepth int
	isLeftBalance := checkBalance(root.Left, &leftDepth)
	isRightBalance := checkBalance(root.Right, &rightDepth)
	if isLeftBalance && isRightBalance {
		if math.Abs(float64(leftDepth-rightDepth)) > 1 {
			return false
		}
		*depth = 1 + int(math.Max(float64(leftDepth), float64(rightDepth)))
		return true
	}
	return false
}

func isBalanced(root *TreeNode) bool {
	var depth int = 0
	return checkBalance(root, &depth)
}
