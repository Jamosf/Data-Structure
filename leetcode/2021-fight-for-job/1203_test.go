package ojeveryday

import "container/list"

// 二叉树的前序遍历迭代遍历
func preorderTraversal_(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	out := make([]int, 0)
	l := list.New()
	l.PushBack(root)
	for l.Len() != 0 {
		v := l.Back()
		l.Remove(v)
		r := v.Value.(*TreeNode)
		out = append(out, r.Val)
		if r.Right != nil {
			l.PushBack(r.Right)
		}
		if r.Left != nil {
			l.PushBack(r.Left)
		}
	}
	return out
}
