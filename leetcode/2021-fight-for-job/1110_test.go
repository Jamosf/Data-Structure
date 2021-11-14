// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
	"testing"
)

// leetcode274: 排序
func hIndex274(citations []int) int {
	n := len(citations)
	sort.Ints(citations)
	for i := 0; i < n; i++ {
		if citations[i] >= n-i {
			return n - i
		}
	}
	return 0
}

func Test_hIndex274(t *testing.T) {
	fmt.Println(hIndex274([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex274([]int{1, 3, 1}))
	fmt.Println(hIndex274([]int{1, 1, 1}))
}

// leetcode324: 排序 朴素解法
func wiggleSort(nums []int) {
	n := len(nums)
	sort.Ints(nums)
	s := make([]int, n)
	n1 := nums[:(n+1)/2]
	n2 := nums[(n+1)/2:]
	for i := 0; i < len(n1); i++ {
		s[i*2] = n1[len(n1)-1-i]
	}
	for i := 0; i < len(n2); i++ {
		s[i*2+1] = n2[len(n2)-1-i]
	}
	copy(nums, s)
}

// leetcode324: 优化解法, 快速选择+3-way-partition, 实现暂时有点问题。
func wiggleSort_(nums []int) {
	n := len(nums)
	mid := quickSort1(nums, 0, n, (n+1)/2)[n/2]
	// 3-way-partition
	i, j, k := 0, 0, n-1
	for j < k {
		if nums[j] > mid {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		} else if nums[i] < mid {
			nums[j], nums[i] = nums[i], nums[j]
			i++
			j++
		} else {
			j++
		}
	}
	n1 := nums[:(n+1)/2]
	n2 := nums[(n+1)/2:]
	s := make([]int, n)
	for i := 0; i < len(n1); i++ {
		s[i*2] = n1[len(n1)-1-i]
	}
	for i := 0; i < len(n2); i++ {
		s[i*2+1] = n2[len(n2)-1-i]
	}
	copy(nums, s)
}

// leetcode324: 桶排序
func wiggleSort__(nums []int) {
	n := len(nums)
	bucket := make([]int, 5001)
	for i := range nums {
		bucket[nums[i]]++
	}
	s := make([]int, n)
	idx := 5000
	// 先安排大的，从后向前可以避免重复的数字挤在一起
	for i := 0; i < n/2; i++ {
		for bucket[idx] == 0 {
			idx--
		}
		s[i*2+1] = idx
		bucket[idx]--
	}
	for i := 0; i < (n+1)/2; i++ {
		for bucket[idx] == 0 {
			idx--
		}
		s[i*2] = idx
		bucket[idx]--
	}
	copy(nums, s)
}

func Test_wiggleSort(t *testing.T) {
	v := []int{3, 2, 3, 3, 2, 1, 1, 2, 3, 1, 1, 3, 2, 1, 2, 2, 2, 2, 1}
	wiggleSort__(v)
	fmt.Println(v)
	v1 := []int{1, 3, 2, 2, 3, 1, 6}
	wiggleSort__(v1)
	fmt.Println(v1)
	v2 := []int{4, 5, 5, 6}
	wiggleSort__(v2)
	fmt.Println(v2)
}
