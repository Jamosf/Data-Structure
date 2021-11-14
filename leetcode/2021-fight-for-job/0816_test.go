// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// 前缀和
// 第一题
// leetcode525: 连续数组
// 前缀和+hash
func findMaxLength(nums []int) int {
	sum := 0
	m := make(map[int]int)
	maxn := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		v := 2*sum - i - 1
		if v == 0 {
			maxn = max(maxn, i+1)
		}
		if idx, ok := m[v]; ok {
			maxn = max(maxn, i-idx)
		} else {
			m[v] = i
		}
	}
	return maxn
}

func Test_findMaxLength(t *testing.T) {
	fmt.Println(findMaxLength([]int{0, 1, 0, 1}))
}

// 第二题

func subarraySum(nums []int, k int) int {
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	cnt := 0
	for i := 0; i < len(sum); i++ {
		for j := i + 1; j < len(sum); j++ {
			if sum[j]-sum[i] == k {
				cnt++
			}
		}
	}
	return cnt
}

func Test_subarraySum(t *testing.T) {
	fmt.Println(subarraySum([]int{-1, -1, 1}, 0))
}

// 第三题
// leetcode1109: 航班预定统计
func corpFlightBookings(bookings [][]int, n int) []int {
	sum := make([]int, n+1)
	sum[0] = 0
	for _, booking := range bookings {
		for i := booking[0]; i <= booking[1]; i++ {
			sum[i] += booking[2]
		}
	}
	return sum[1:]
}

// leetcode1109: 航班预定统计
// 差分数组+前缀和
func corpFlightBookings_(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	for _, booking := range bookings {
		diff[booking[0]-1] += booking[2]
		diff[booking[1]] -= booking[2]
	}
	for i := 1; i < len(diff); i++ {
		diff[i] += diff[i-1]
	}
	return diff[:n]
}

// 并查集
// 第四题
// leetcode128: 最长连续序列
func longestConsecutive(nums []int) int {
	inf := int(1e9 + 1)
	m := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = struct{}{}
	}
	maxn := 0
	for k := 0; k < len(nums); k++ {
		if _, ok := m[nums[k]]; !ok {
			continue
		}
		pre := nums[k]
		i, j := pre, pre
		for ; i < inf; i++ {
			if _, ok := m[i+1]; !ok {
				break
			}
			delete(m, i+1)
		}
		for ; j > -1*inf; j-- {
			if _, ok := m[j-1]; !ok {
				break
			}
			delete(m, j-1)
		}
		delete(m, pre)
		fmt.Println(i, j)
		maxn = max(maxn, i-j+1)
	}

	return maxn
}

func Test_longestConsecutive(t *testing.T) {
	fmt.Println(longestConsecutive([]int{4, 0, -4, -2, 2, 5, 2, 0, -8, -8, -8, -8, -1, 7, 4, 5, 5, -4, 6, 6, -3}))
}
