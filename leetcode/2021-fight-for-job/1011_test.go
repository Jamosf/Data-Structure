// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
	"testing"
)

// 第一题
// O(n^2)解法
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
		for j := i - 1; j >= 0; j-- {
			if nums[j] >= i-j {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func Test_jump(t *testing.T) {
	fmt.Println(jump1([]int{1, 1, 1, 4, 1, 1, 1}))
}

// 第二题
// O(n)解法
func jump1(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum := 0
	maxn := math.MinInt32
	idx := -1
	for i := n - 1; i >= 0; i-- {
		sum += gas[i] - cost[i]
		if sum > maxn {
			maxn = sum
			idx = i
		}
	}
	return idx
}

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		s, t := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		return s+t > t+s
	})
	var buff bytes.Buffer
	for i := range nums {
		if buff.Len() == 0 && nums[i] == 0 {
			continue
		}
		buff.WriteString(strconv.Itoa(nums[i]))
	}
	return buff.String()
}
