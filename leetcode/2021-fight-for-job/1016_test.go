// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

// tag-[二叉树]
// leetcode106, 可以使用map优化查找节点的idx
func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	n := len(postorder)
	root := &TreeNode{Val: postorder[n-1]}
	idx := n - 1
	for idx >= 0 && inorder[idx] != postorder[n-1] {
		idx--
	}
	root.Right = buildTree1(inorder[idx+1:], postorder[idx:n-1])
	root.Left = buildTree1(inorder[:idx], postorder[:idx])
	return root
}

// tag-[线段树]
// leetcode307 线段树
type NumArray1 struct {
	tree *segTree
}

type segTree struct {
	start, end int
	sum        int
	left       *segTree
	right      *segTree
}

func buildSegTree(start, end int, nums []int) *segTree {
	if start == end {
		return &segTree{start: start, end: end, sum: nums[start]}
	}
	mid := (start + end) >> 1
	left := buildSegTree(start, mid, nums)
	right := buildSegTree(mid+1, end, nums)
	return &segTree{start: start, end: end, sum: left.sum + right.sum, left: left, right: right}
}

func update(root *segTree, index int, val int) {
	if root.start == root.end && root.start == index {
		root.sum = val
		return
	}
	mid := (root.start + root.end) >> 1
	if index <= mid {
		update(root.left, index, val)
	} else {
		update(root.right, index, val)
	}
	root.sum = root.left.sum + root.right.sum
}

func query(root *segTree, i, j int) int {
	if root.start == i && root.end == j {
		return root.sum
	}
	mid := (root.start + root.end) >> 1
	if j <= mid {
		return query(root.left, i, j)
	} else if i > mid {
		return query(root.right, i, j)
	} else {
		return query(root.left, i, mid) + query(root.right, mid+1, j)
	}
}

func ConstructorNumArray1(nums []int) NumArray1 {
	return NumArray1{tree: buildSegTree(0, len(nums)-1, nums)}
}

func (n *NumArray1) Update(index int, val int) {
	update(n.tree, index, val)
}

func (n *NumArray1) SumRange(left int, right int) int {
	return query(n.tree, left, right)
}
