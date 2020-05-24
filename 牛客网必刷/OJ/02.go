package main

// 状态方程: dp[n] = 1 , dp[n] = dp[n-1] - (dp[n-1]/3+1)
func cakeNumber(n int) int {
	// write code here
	dp := make([]int, n)
	dp[n-1] = 1
	for i := n - 1; i > 0; i-- {
		dp[i-1] = (dp[i]*3 + 3) / 2
	}
	return dp[0]
}

// 状态方程:
func solve(s string) int {
	// write code here
	return 1
}
