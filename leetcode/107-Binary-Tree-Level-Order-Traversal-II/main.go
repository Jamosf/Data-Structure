package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(s *[][]int, level int, root *TreeNode) {
	if root == nil {
		return
	}
	if len(*s) == level {
		*s = append(*s, []int{})
	}
	(*s)[level] = append((*s)[level], root.Val)
	for _, r := range []*TreeNode{root.Left, root.Right} {
		dfs(s, level+1, r)
	}
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var s [][]int
	dfs(&s, 0, root)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}
