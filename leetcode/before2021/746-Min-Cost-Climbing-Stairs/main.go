package main

import (
	"fmt"
	"math"
)

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n < 1 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < n-1; i++ {
		dp[i] = int(math.Min(float64(dp[i-1]), float64(dp[i-2]))) + cost[i]
	}
	dp[n-1] = int(math.Min(float64(dp[n-2]), float64(dp[n-3]+cost[n-1])))
	return dp[n-1]
}

func main() {
	test := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(test))
	test = []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(test))
}
