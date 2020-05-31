package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			return true
		}
	}
	if root.Left != nil {
		if hasPathSum(root.Left, sum) {
			return true
		}
	}
	if root.Right != nil {
		if hasPathSum(root.Right, sum) {
			return true
		}
	}
	return false
}
