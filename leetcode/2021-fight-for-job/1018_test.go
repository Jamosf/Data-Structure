// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// leetcode673: 可以用双dp来理解
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	maxn := 1
	count := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
		if dp[i] > maxn {
			maxn = dp[i]
			ans = count[i]
		} else if dp[i] == maxn {
			ans += count[i]
		}
	}
	return ans
}

func Test_findNumberOfLIS(t *testing.T) {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	fmt.Println(findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2}))
}
