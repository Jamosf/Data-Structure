// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// 栈的解法，栈保留索引
func longestValidParentheses(s string) int {
	n := len(s)
	stack := make([]int, 0)
	cnt := 0
	maxn := 0
	for i := 0; i < n; i++ {
		if len(stack) != 0 && isPair(s[stack[len(stack)-1]], s[i]) {
			stack = stack[:len(stack)-1]
			cnt += 2
		} else {
			if s[i] == ')' {
				cnt = 0
				stack = stack[:0]
			}
			stack = append(stack, i)
		}
		if len(stack) != 0 && s[stack[len(stack)-1]] == '(' {
			maxn = max(maxn, i-stack[len(stack)-1])
		} else {
			maxn = max(maxn, cnt)
		}
	}
	return maxn
}

func isPair(a, b byte) bool {
	return a == '(' && b == ')'
}

func Test_longestValidParentheses(t *testing.T) {
	fmt.Println(longestValidParentheses1("()()(((()())))"))
}

// 动态规划解法
func longestValidParentheses1(s string) int {
	n := len(s)
	dp := make([]int, n) // 表示以）结尾的最长的子串长度
	dp[0] = 0
	maxn := 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-2*dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 1
			}
		}
		// if s[x] == '(' {
		// 	dp[x] = dp[x-1]
		// }
		maxn = max(maxn, dp[i])
	}
	return 2 * maxn
}

func trap(height []int) int {
	n := len(height)
	maxn := height[0]
	for i := 1; i < n; i++ {
		maxn = max(maxn, height[i])
	}
	sum := 0
	for i := 0; i < maxn; i++ {
		hasHighInLeft := false
		tmp := 0
		for j := 0; j < n; j++ {
			if hasHighInLeft && height[j] < i {
				tmp++
			}
			if height[j] >= i {
				sum += tmp
				tmp = 0
				hasHighInLeft = true
			}
		}
	}
	return sum
}