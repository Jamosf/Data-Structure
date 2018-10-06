package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, level int, average *[]float64, count *[]int) {
	if root == nil {
		return
	}
	if len(*average) == level {
		*average = append(*average, float64(0))
		*count = append(*count, 0)
	}
	(*average)[level] += float64(root.Val)
	(*count)[level] += 1
	for _, r := range []*TreeNode{root.Left, root.Right} {
		dfs(r, level+1, average, count)
	}
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var r []float64
	var count []int
	dfs(root, 0, &r, &count)
	for i := 0; i < len(r); i++ {
		r[i] = r[i] / float64(count[i])
	}
	return r
}
