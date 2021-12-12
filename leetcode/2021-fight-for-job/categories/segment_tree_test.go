package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

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