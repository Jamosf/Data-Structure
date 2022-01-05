package ojeveryday

// tag-[动态规划/字符串]
// leetcode516: 最长回文子序列
// dp[i][j]表示i...j范围内回文子序列的长度
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

// tag-[动态规划]
// leetcode123: 买卖股票的最佳时机 III
// dp[i][k][j], k表示交易次数
func maxProfitIII(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	n := len(prices)
	dp := make([][3][2]int, n)

	for i := 0; i < n; i++ {
		for k := 2; k >= 1; k-- {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[0]
			} else {
				dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
				dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
			}
		}
	}
	return dp[n-1][2][0]
}

// tag-[动态规划]
// leetcode188: 买卖股票的最佳时机IV
// dp[i][k][j], k表示交易次数
func maxProfitIV(k int, prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	n := len(prices)
	dp := make([][][2]int, n)
	for i := range dp {
		dp[i] = make([][2]int, k+1)
	}

	for i := 0; i < n; i++ {
		for j := k; j >= 1; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[0]
			} else {
				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
				dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
			}
		}
	}
	return dp[n-1][k][0]
}

// tag-[动态规划]
// leetcode714: 买卖股票的最佳时机含手续费
// dp[i][k][j], k表示交易次数
func maxProfitV(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0] - fee
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}
	return dp[n-1][0]
}
