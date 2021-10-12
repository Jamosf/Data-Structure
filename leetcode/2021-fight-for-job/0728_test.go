// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"testing"
)

// 第一题
func isPowerOfTwo(n int) bool {
	cnt := 0
	for n != 0 {
		n &= n - 1
		cnt++
	}
	return cnt == 1
}

// 第二题
func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num &= num - 1
		cnt++
	}
	return cnt
}

// 第三题
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

// 第四题
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

// 第五题
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
