package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, sum int, count *int) {
	if root == nil {
		return
	}
	sum -= root.Val
	if sum == 0 {
		*count = *count + 1
	}
	dfs(root.Left, sum, count)
	dfs(root.Right, sum, count)
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var count int
	dfs(root, sum, &count)
	return count + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
