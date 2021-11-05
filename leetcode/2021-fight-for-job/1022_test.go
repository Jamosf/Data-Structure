// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// leetcode63:动态规划(可以降维)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}

// leetcode95:回溯法
func generateTrees(n int) []*TreeNode {
	var backtrace func(start, end int) []*TreeNode
	backtrace = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		allTrees := make([]*TreeNode, 0)
		for i := start; i <= end; i++ {
			leftTrees := backtrace(start, i-1)
			rightTrees := backtrace(i+1, end)
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					currTree := &TreeNode{i, nil, nil}
					currTree.Left = left
					currTree.Right = right
					allTrees = append(allTrees, currTree)
				}
			}
		}
		return allTrees
	}
	return backtrace(1, n)
}

// leetcode131: 分割回文串(典型的回溯)
func partition(s string) [][]string {
	n := len(s)
	ans := make([][]string, 0)
	tmp := make([]string, 0)
	isPlalindrome := func(b string) bool {
		j := len(b) - 1
		for i := 0; i < j; i++ {
			if b[i] != b[j] {
				return false
			}
			j--
		}
		return true
	}

	var backtrace func(idx int)
	backtrace = func(idx int) {
		if idx == n {
			ans = append(ans, append([]string{}, tmp...))
			return
		}
		for i := idx + 1; i <= n; i++ {
			if isPlalindrome(s[idx:i]) {
				tmp = append(tmp, s[idx:i])
				backtrace(i)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtrace(0)
	return ans
}

// leetcode131: 分割回文串(典型的回溯) 使用记忆化优化或者使用dp来预处理字符串把任意i->j是否为回文算出来
func partition_(s string) [][]string {
	n := len(s)
	ans := make([][]string, 0)
	tmp := make([]string, 0)
	dp := make([][]int8, n)
	for i := range dp {
		dp[i] = make([]int8, n)
	}
	var isPlalindrome func(i, j int) int8
	isPlalindrome = func(i, j int) int8 {
		if i >= j {
			return 1
		}
		if dp[i][j] != 0 {
			return dp[i][j]
		}
		dp[i][j] = -1
		if s[i] == s[j] {
			dp[i][j] = isPlalindrome(i+1, j-1)
		}
		return dp[i][j]
	}

	var backtrace func(idx int)
	backtrace = func(idx int) {
		if idx == n {
			ans = append(ans, append([]string{}, tmp...))
			return
		}
		for i := idx; i < n; i++ {
			if isPlalindrome(idx, i) > 0 {
				tmp = append(tmp, s[idx:i+1])
				backtrace(i + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtrace(0)
	return ans
}

func Test_partition(t *testing.T) {
	fmt.Println(partition_("aabbaababab"))
}

// leetcode229: 求众数
func majorityElement229(nums []int) []int {
	n := len(nums)
	e1, e2 := 0, 0
	vote1, vote2 := 0, 0
	for i := range nums {
		if vote1 > 0 && nums[i] == e1 {
			vote1++
		} else if vote2 > 0 && nums[i] == e2 {
			vote2++
		} else if vote1 == 0 {
			e1 = nums[i]
			vote1++
		} else if vote2 == 0 {
			e2 = nums[i]
			vote2++
		} else {
			vote1--
			vote2--
		}
	}
	// 验证
	cnt1, cnt2 := 0, 0
	for i := range nums {
		if vote1 > 0 && nums[i] == e1 {
			cnt1++
		} else if vote2 > 0 && nums[i] == e2 {
			cnt2++
		}
	}
	var ans []int
	if vote1 > 0 && cnt1*3 > n {
		ans = append(ans, e1)
	}
	if vote2 > 0 && cnt2*3 > n {
		ans = append(ans, e2)
	}
	return ans
}
