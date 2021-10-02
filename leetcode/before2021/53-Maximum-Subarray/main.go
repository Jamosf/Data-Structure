package main

import "fmt"

//dp[i] = dp[i-1] > 0 ? dp[i-1] : A[i]
func maxSubArray(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func main() {
	test := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(test))
}
