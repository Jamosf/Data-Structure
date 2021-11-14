// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// leetcode974:前缀和
func subarraysDivByK(nums []int, k int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	cnt := make(map[int]int)
	for i := range nums {
		cnt[(nums[i]%k+k)%k]++
	}
	ans := 0
	for i := range cnt {
		if i == 0 {
			ans += cnt[i]
		}
		ans += cnt[i] * (cnt[i] - 1) / 2
	}
	return ans
}

func Test_subarraysDivByK(t *testing.T) {
	// fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
	// fmt.Println(subarraysDivByK([]int{-1, 2, 9}, 2))
	fmt.Println(subarraysDivByK([]int{-6, 6}, 5))
}

// leetcode1829:前缀异或
func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	xor := make([]int, n)
	xor[0] = nums[0]
	for i := 1; i < n; i++ {
		xor[i] = xor[i-1] ^ nums[i]
	}
	ans := make([]int, 0, n)
	maxn := 1<<maximumBit - 1
	for i := n - 1; i >= 0; i-- {
		ans = append(ans, maxn^xor[i])
	}
	return ans
}

func Test_getMaximumXor(t *testing.T) {
	fmt.Println(getMaximumXor([]int{0, 1, 1, 3}, 2))
	fmt.Println(getMaximumXor([]int{2, 3, 4, 7}, 3))
	fmt.Println(getMaximumXor([]int{0, 1, 2, 2, 5, 7}, 3))
}

// leetcode769: 数组分组，前k个的最大值是不是k
func maxChunksToSorted(arr []int) int {
	n := len(arr)
	maxn := arr[0]
	ans := 0
	for i := 0; i < n; i++ {
		maxn = max(maxn, arr[i])
		if maxn == i {
			ans++
		}
	}
	return ans
}

func Test_maxChunksToSorted(t *testing.T) {
	fmt.Println(maxChunksToSorted([]int{4, 3, 2, 1, 0}))
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
	fmt.Println(maxChunksToSorted([]int{2, 0, 1}))
}
