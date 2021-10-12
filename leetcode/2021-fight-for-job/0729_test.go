// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"testing"
)

// 第一题
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 第二题
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	if root == nil {
		return true
	}
	if !isValidBST(root.Left) {
		return false
	}
	if root.Val <= pre {
		return false
	}
	pre = root.Val
	return isValidBST(root.Right)
}

// 方法2
func isValidBST2(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val >= upper || root.Val <= lower {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

// 第三题
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	return dfs(root, m, k)
}

func dfs(root *TreeNode, m map[int]struct{}, k int) bool {
	if root == nil {
		return false
	}
	if _, ok := m[root.Val]; ok {
		return true
	}
	m[k-root.Val] = struct{}{}
	return dfs(root.Left, m, k) || dfs(root.Right, m, k)
}

// 第四题
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		if left := lowestCommonAncestor(root.Left, p, q); left != nil {
			return left
		}
	}
	if root.Val < p.Val && root.Val < q.Val {
		if right := lowestCommonAncestor(root.Right, p, q); right != nil {
			return right
		}
	}
	return root
}

// 第五题
func reverseBits(num uint32) uint32 {
	var ret uint32
	for i := 0; i < 32; i++ {
		bit := (num >> i) & 1
		bit <<= 31 - i
		ret += bit
	}
	return ret
}

// 方法2
func reverseBits1(n uint32) uint32 {
	n = (n >> 16) | (n << 16)
	n = ((n & 0xff00ff00) >> 8) | ((n & 0x00ff00ff) << 8)
	n = ((n & 0xf0f0f0f0) >> 4) | ((n & 0x0f0f0f0f) << 4)
	n = ((n & 0xcccccccc) >> 2) | ((n & 0x33333333) << 2)
	n = ((n & 0xaaaaaaaa) >> 1) | ((n & 0x55555555) << 1)
	return n
}

func Test_reverseBits(t *testing.T) {
	fmt.Println(reverseBits(0b00000010100101000001111010011100))
}

// 第六题
func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		ret ^= nums[i]
	}
	return ret
}

// 第七题
func permute(nums []int) [][]int {
	n := len(nums)
	l := factorial(n)
	res := make([][]int, l, l)
	for i := 0; i < l; i++ {
		res[i] = make([]int, n, n)
	}
	idx := 0
	backtrack(res, nums, 0, n, &idx)
	return res
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func backtrack(res [][]int, nums []int, first, len int, idx *int) {
	if first == len {
		copy(res[*idx], nums)
		*idx++
		return
	}
	for i := first; i < len; i++ {
		nums[i], nums[first] = nums[first], nums[i]
		backtrack(res, nums, first+1, len, idx)
		nums[i], nums[first] = nums[first], nums[i]
	}
}

func Test_permute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}
