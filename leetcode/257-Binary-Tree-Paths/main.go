package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, s *string, r *[]string) {
	if root == nil {
		*s = ""
		return
	}
	if root.Left != nil || root.Right != nil {
		*s += strconv.Itoa(root.Val) + "->"
		var left, right string = *s, *s
		if root.Left != nil {
			dfs(root.Left, &left, r)
		}
		if root.Right != nil {
			dfs(root.Right, &right, r)
		}
	}
	if root.Left == nil && root.Right == nil {
		*s += strconv.Itoa(root.Val)
		*r = append(*r, *s)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var s string
	var r []string
	dfs(root, &s, &r)
	return r
}
