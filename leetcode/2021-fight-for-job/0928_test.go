// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
	"testing"
)

func sortColors(nums []int) {
	n := len(nums)
	left, right := 0, n-1
	for i := 0; i < right; i++ {
		for ; i <= right && nums[i] == 2; right-- {
			nums[right], nums[i] = nums[i], nums[right]
		}
		if nums[i] == 0 && i >= left {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
}

func Test_sortColors(t *testing.T) {
	sortColors([]int{1, 1, 1, 1, 2, 2, 2, 2, 0, 0, 0})
}

func maxProduct1(nums []int) int {
	n := len(nums)
	dp := make([][2]int, n) // 以i结尾的最大值或最小值，
	dp[0][0] = nums[0]      // 0存最小值
	dp[0][1] = nums[0]      // 1存最大值
	maxn := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] >= 0 {
			dp[i][0] = min(dp[i-1][0]*nums[i], nums[i])
			dp[i][1] = max(dp[i-1][1]*nums[i], nums[i])
		} else {
			dp[i][0] = min(dp[i-1][1]*nums[i], nums[i])
			dp[i][1] = max(dp[i-1][0]*nums[i], nums[i])
		}
		maxn = max(maxn, dp[i][1])
	}
	return maxn
}

func Test_maxProduct1(t *testing.T) {
	fmt.Println(maxProduct1([]int{-3, 2, -4}))
}

func massage(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	if n < 2 {
		return dp[0]
	}
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		for j := i - 2; j >= 0; j-- {
			dp[i] = max(dp[i], dp[j]+nums[i])
		}
	}
	return max(dp[n-1], dp[n-2])
}

// 树形dp
func rob1(root *TreeNode) int {
	var dfs func(r *TreeNode) [2]int
	dfs = func(r *TreeNode) [2]int {
		if r == nil {
			return [2]int{}
		}
		L := dfs(r.Left)
		R := dfs(r.Right)

		dp := [2]int{} // 0表示不选父节点、1表示选择父节点
		dp[0] = max(L[0], L[1]) + max(R[0], R[1])
		dp[1] = r.Val + L[0] + R[0]
		return dp
	}
	ans := dfs(root)
	return max(ans[0], ans[1])
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1]
			}
			if j == 0 && i > 0 {
				dp[i][j] = dp[i-1][j]
			}
			if i > 0 && j > 0 {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}

func sortList(head *ListNode) *ListNode {
	p := head
	t := make([]int, 0)
	for p != nil {
		t = append(t, p.Val)
		p = p.Next
	}
	sort.Ints(t)
	out := &ListNode{}
	q := out
	for i := 0; i < len(t); i++ {
		out.Next = &ListNode{Val: t[i]}
		out = out.Next
	}
	return q.Next
}

// 递归解法
func sortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	tmp := slow.Next
	left, right := sortList1(head), sortList1(tmp)
	t := &ListNode{}
	p := t
	for left != nil && right != nil {
		if left.Val < right.Val {
			t.Next = left
			left = left.Next
		} else {
			t.Next = right
			right = right.Next
		}
	}
	if left == nil {
		t.Next = right
	} else {
		t.Next = left
	}
	return p.Next
}

// 每个数字作为根节点，左右子树的各种组合乘积
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
