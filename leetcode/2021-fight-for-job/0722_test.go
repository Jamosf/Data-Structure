// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 第一题
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	m := &TreeNode{}
	m.Val = root1.Val + root2.Val
	m.Left = mergeTrees(root1.Left, root2.Left)
	m.Right = mergeTrees(root1.Right, root2.Right)

	return m
}

// 第二题

// type Node struct {
// 	Val   int
// 	Left  *Node
// 	Right *Node
// 	Next  *Node
// }

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	return nil
}
