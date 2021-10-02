package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil || root.Right != nil {
		root.Left, root.Right = root.Right, root.Left
	}
	if root.Left != nil {
		dfs(root.Left)
	}
	if root.Right != nil {
		dfs(root.Right)
	}
}

func invertTree(root *TreeNode) *TreeNode {
	dfs(root)
	return root
}
