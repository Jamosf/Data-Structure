// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"testing"
)

// tag-[数学]
// 第一题
// leetcode191: 2的幂
func isPowerOfTwo(n int) bool {
	cnt := 0
	for n != 0 {
		n &= n - 1
		cnt++
	}
	return cnt == 1
}

// tag-[数学]
// 第二题
// leetcode461: 汉明距离
func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num &= num - 1
		cnt++
	}
	return cnt
}

// tag-[动态规划]
// 第三题
// leetcode198: 打家劫舍
func rob(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else if i == 1 {
			dp[i] = max(nums[0], nums[1])
		} else {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}
	}
	return dp[len(nums)-1]
}

// tag-[动态规划]
// 第四题
// leetcode120: 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	m := len(triangle)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	minn := math.MaxInt64
	for j := 0; j < len(dp[m-1]); j++ {
		minn = min(minn, dp[m-1][j])
	}
	return minn
}

func Test_minimumTotal(t *testing.T) {
	fmt.Println(minimumTotal([][]int{{-1}, {-2, -3}}))
}

// tag-[二叉树]
// 第五题
// leetcode700: 二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}
