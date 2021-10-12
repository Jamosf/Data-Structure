// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"testing"
)

// 第一题
// 背包问题
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	diff := sum - target
	if diff < 0 || diff%2 != 0 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := neg; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[neg]
}

func Test_findTargetSumWays(t *testing.T) {
	fmt.Println(findTargetSumWays([]int{1}, 1))
}

// 第二题
// 完全平方数，背包问题
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func Test_numSquares(t *testing.T) {
	fmt.Println(numSquares(12))
}

// 第三题
// 组合总和IV
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 1; j <= target; j++ { // 背包
		for i := 1; i <= len(nums); i++ { // 物品
			if nums[i-1] <= j {
				dp[j] += dp[j-nums[i-1]]
			}
		}
	}
	return dp[target]
}

func Test_combinationSum4(t *testing.T) {
	fmt.Println(combinationSum4([]int{9}, 4))
}

// 第四题
// 零钱兑换II
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 1; i <= len(coins); i++ { // 物品
		for j := coins[i-1]; j <= amount; j++ { // 背包
			dp[j] += dp[j-coins[i-1]]
		}
	}
	return dp[amount]
}

// 第五题
// 一和零
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	getOne := func(s string) int {
		cnt := 0
		for i := range s {
			if s[i] == '1' {
				cnt++
			}
		}
		return cnt
	}
	type zeroOne struct {
		zero int
		one  int
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	strsInfo := make([]zeroOne, len(strs))
	for i := range strs {
		one := getOne(strs[i])
		strsInfo[i] = zeroOne{zero: len(strs[i]) - one, one: one}
	}
	for i := 1; i <= len(strsInfo); i++ { // 物品
		for j := m; j >= strsInfo[i-1].zero; j-- { // 背包
			for k := n; k >= strsInfo[i-1].one; k-- { // 背包
				dp[j][k] = max(dp[j][k], dp[j-strsInfo[i-1].zero][k-strsInfo[i-1].one]+1)
			}
		}
	}
	return dp[m][n]
}

func Test_findMaxForm(t *testing.T) {
	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
}

// 第六题
// 最后一块石头的重量 II
func lastStoneWeightII(stones []int) int {
	sum := 0
	for i := range stones {
		sum += stones[i]
	}
	neg := sum / 2
	dp := make([]bool, neg+1)
	dp[0] = true
	for i := 1; i <= len(stones); i++ { // 物品
		for j := neg; j >= stones[i-1]; j-- { // 背包
			dp[j] = dp[j] || dp[j-stones[i-1]]
		}
	}
	maxn := 0
	for i := range dp {
		if dp[i] {
			maxn = i
		}
	}
	return sum - 2*maxn
}
