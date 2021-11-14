// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// leetcode46: 全排列
func permute(nums []int) [][]int {
	n := len(nums)
	l := factorial(n)
	ans := make([][]int, 0, l)
	var backtracking func(level int)
	backtracking = func(level int) {
		if level == n {
			t := make([]int, n)
			copy(t, nums)
			ans = append(ans, t)
		}
		for i := level; i < n; i++ {
			nums[i], nums[level] = nums[level], nums[i]
			backtracking(level + 1)
			nums[i], nums[level] = nums[level], nums[i]
		}
	}
	backtracking(0)
	return ans
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func Test_permute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}

// leetcode77: 组合
func combine(n int, k int) [][]int {
	var ans [][]int
	tmp := make([]int, 0, k)
	var backtracking func(level int)
	backtracking = func(idx int) {
		if len(tmp) == k {
			t := make([]int, k)
			copy(t, tmp)
			ans = append(ans, t)
		}
		for i := idx; i <= n; i++ {
			if len(tmp)+(n-i+1) >= k {
				tmp = append(tmp, i)
				backtracking(i + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtracking(1)
	return ans
}

func Test_combine(t *testing.T) {
	fmt.Println(combine(4, 3))
}
