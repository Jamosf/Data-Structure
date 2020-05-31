package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode, k int, flag *bool, help map[int]bool) {
	if root == nil {
		return
	}
	traverse(root.Left, k, flag, help)
	traverse(root.Right, k, flag, help)
	if _, ok := help[(k - root.Val)]; ok {
		*flag = true
		return
	}
	help[root.Val] = true
}

func findTarget(root *TreeNode, k int) bool {
	var flag bool
	help := make(map[int]bool)
	traverse(root, k, &flag, help)
	return flag
}
