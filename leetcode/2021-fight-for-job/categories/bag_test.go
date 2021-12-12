package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[背包]
// 第七题
// leetcode 剑指offer14-I: 剪绳子
// 剪绳子类似于整数拆分：dp[i] = max(dp[j]*(i-j), j*(i, j))
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	mod := int(1e9 + 7)
	for i := 2; i <= n; i++ {
		for j := i - 1; j > 0; j-- {
			dp[i] = max(dp[i]%mod, (dp[j]%mod)*(i-j)%mod)
			dp[i] = max(dp[i]%mod, (j%mod)*(i-j)%mod)
		}
	}
	return dp[n] % mod
}

// tag-[背包]
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

// tag-[背包]
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

// tag-[背包]
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
// tag-[背包]
// 第一题
// leetcode494: 目标和
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

// tag-[背包]
// 第二题
// leetcode279: 完全平方数，背包问题
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

// tag-[背包]
// 第三题
// leetcode377: 组合总和IV
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

// tag-[背包]
// 第四题
// leetcode518: 零钱兑换II
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

// tag-[背包]
// 第五题
// leetcode474: 一和零
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

// tag-[背包]
// 第六题
// leetcode1049: 最后一块石头的重量 II
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
// tag-[背包]
// 第三题
// TODO
func combinationSum1(candidates []int, target int) [][]int {
	res := make([][][]int, target+1)
	res[0] = make([][]int, 0)
	for i := 1; i < len(candidates); i++ { // 物品
		for j := candidates[i-1]; j <= target; j++ { // 背包
			for k := 0; k < len(res[j-candidates[i-1]]); k++ {
				res[j-candidates[i-1]][k] = append(res[j-candidates[i-1]][k], candidates[i])
			}
			res[j] = append(res[j], res[j-candidates[i-1]]...)
		}
	}
	return res[target]
}

func Test_combinationSum1(t *testing.T) {
	fmt.Println(combinationSum1([]int{2, 3, 6, 7}, 7))
}