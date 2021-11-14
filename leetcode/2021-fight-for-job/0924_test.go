// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
	"testing"
)

// 第一题
// leetcode1806: 还原排列的最少操作步数
// 位运算
func reinitializePermutation(n int) int {
	ans := 0
	i := 1
	for {
		if i&1 == 0 {
			i = i >> 1
		} else {
			i = n>>1 + (i-1)>>1
		}
		ans++
		if i == 1 {
			break
		}
	}
	return ans
}

func Test_reinitializePermutation(t *testing.T) {
	fmt.Println(reinitializePermutation(8))
}

// 第二题
// leetcode22: 括号生成
func generateParenthesis(n int) []string {
	s := make([]byte, 2*n)
	for i := 0; i < 2*n; i++ {
		if i < n {
			s[i] = '('
		} else {
			s[i] = ')'
		}
	}
	ans := make([]string, 0)
	t := make([]byte, 2*n)
	var backtracking func(lvl int)
	backtracking = func(first int) {
		if first == n {
			if isValid(s) && notContain(string(s), ans) {
				copy(t, s)
				ans = append(ans, string(t))
			}
			return
		}
		for i := first; i < 2*n; i++ {
			if i == first || s[i] != s[first] {
				s[i], s[first] = s[first], s[i]
				backtracking(first + 1)
				s[i], s[first] = s[first], s[i]
			}
		}
	}
	backtracking(0)
	return ans
}

func isValid(s []byte) bool {
	cnt := 0
	for i := range s {
		if s[i] == '(' {
			cnt++
		} else {
			cnt--
		}
		if cnt < 0 {
			return false
		}
	}
	return cnt == 0
}

func notContain(s string, t []string) bool {
	for i := range t {
		if t[i] == s {
			return false
		}
	}
	return true
}

func Test_generateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis(8))
}

// 第三题
// leetcode31: 下一个排列
func nextPermutation(nums []int) {
	n := len(nums)
	left, right := -1, -1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			left = i
			break
		}
	}
	if left >= 0 {
		for i := n - 1; i > left; i-- {
			if nums[i] > nums[left] {
				right = i
				break
			}
		}
	}
	if right > 0 {
		nums[left], nums[right] = nums[right], nums[left]
		sort.Ints(nums[left+1:])
		return
	}
	sort.Ints(nums)
}

func Test_nextPermutation(t *testing.T) {
	nums := []int{1, 2, 3, 5, 4}
	nextPermutation(nums)
	fmt.Println(nums)
}

// 第四题
// leetcode621: 任务调度器
func leastInterval(tasks []byte, n int) int {
	f := [26]int{}
	maxn := 0
	// 1. 先找出任务数量最大的
	for i := range tasks {
		v := tasks[i] - 'A'
		f[v]++
		maxn = max(maxn, f[v])
	}
	cnt := 0
	// 2. 计算最后一个桶的数量，只有数量和最大的相同，才有可能占用最后一个桶
	for i := range f {
		if f[i] == maxn {
			cnt++
		}
	}
	// 3. 任务很稀疏时，值为任务数量
	return max(len(tasks), cnt+(n+1)*(maxn-1))
}

// 第五题
// leetcode114: 二叉树展开为链表
func flatten(root *TreeNode) {
	dummy := &TreeNode{}
	p := dummy
	var traval func(r *TreeNode)
	traval = func(r *TreeNode) {
		if r == nil {
			return
		}
		dummy.Right = &TreeNode{Val: r.Val}
		dummy = dummy.Right
		traval(r.Left)
		traval(r.Right)
	}
	traval(root)
	if root == nil {
		return
	}
	root.Left = nil
	root.Right = p.Right.Right
}

func Test_flatten(t *testing.T) {
	flatten(&TreeNode{Right: &TreeNode{1, nil, nil}, Left: &TreeNode{2, nil, nil}, Val: 0})
}

// 第六题
// leetcode55: 跳跃游戏
func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] >= i-j {
				dp[i] = dp[i] || dp[j]
			}
			if dp[i] {
				break
			}
		}
		if !dp[i] {
			return false
		}
	}
	return dp[n-1]
}
