// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import "math"

// 完全背包问题
// leetcode139: 单词拆分
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ { // 遍历背包
		for j := range wordDict { // 遍历物品
			l := len(wordDict[j])
			if i-l >= 0 && s[i-l:i] == wordDict[j] {
				dp[i] = dp[i] || dp[i-l] // 第j个单词是否加入
			}
		}
	}
	return dp[len(s)]
}

// 背包问题
// leetcode416: 分割等和子集
// 0-1背包问题
// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	cap := sum / 2
	dp := make([]bool, cap+1)
	dp[0] = true
	for i := 1; i <= len(nums); i++ {
		for j := cap; j >= nums[i-1]; j-- {
			dp[j] = dp[j] || dp[j-nums[i-1]]
		}
	}
	return dp[cap]
}

// 完全背包问题
// leetcode322: 钱币兑换
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i <= len(coins); i++ {
		for j := coins[i-1]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coins[i-1]]+1)
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
