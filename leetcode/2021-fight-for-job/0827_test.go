// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func findGCD(nums []int) int {
	minn, maxn := nums[0], nums[0]
	for i := range nums {
		if nums[i] > maxn {
			maxn = nums[i]
		}
		if nums[i] < minn {
			minn = nums[i]
		}
	}
	for maxn*minn != 0 && maxn%minn != 0 {
		maxn, minn = minn, maxn%minn
	}
	return minn
}

func Test_findGCD(t *testing.T) {
	fmt.Println(findGCD([]int{1, 12}))
}

func findDifferentBinaryString(nums []string) string {
	m := make(map[string]bool)
	for i := range nums {
		m[nums[i]] = true
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums[i]) - 1; j >= 0; j-- {
			if nums[i][j] == '0' {
				b := []byte(nums[i])
				b[j] = '1'
				if _, ok := m[string(b)]; !ok {
					return string(b)
				}
			}
			if nums[i][j] == '1' {
				b := []byte(nums[i])
				b[j] = '0'
				if _, ok := m[string(b)]; !ok {
					return string(b)
				}
			}
		}
	}
	return ""
}

func Test_findDifferentBinaryString(t *testing.T) {
	fmt.Println(findDifferentBinaryString([]string{"1"}))
}

func minimizeTheDifference(mat [][]int, target int) int {
	m, n := len(mat), len(mat[0])
	for i := range mat {
		sort.Ints(mat[i])
	}
	minn := math.MaxInt32
	var backtrack func(level int)
	var sum int
	var dp [71][4901]bool
	backtrack = func(level int) {
		if sum-target > minn || dp[level][sum] {
			return
		}
		dp[level][sum] = true
		if level == m {
			if abs(sum, target) < minn {
				minn = abs(sum, target)
			}
			return
		}
		for i := 0; i < n; i++ {
			sum += mat[level][i]
			backtrack(level + 1)
			sum -= mat[level][i]
		}
	}
	backtrack(0)
	return minn
}

func Test_min(t *testing.T) {
	fmt.Println(minimizeTheDifference([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 13))
}
