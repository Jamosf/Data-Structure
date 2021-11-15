// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"sort"
)

// 第一题
// leetcode 剑指offer53-I: 在排序数组中查找数字I
func search(nums []int, target int) int {
	idx := sort.Search(len(nums), func(i int) bool {
		return target == nums[i]
	})
	ans := 0
	for i := idx; i < len(nums); i++ {
		if nums[i] == target {
			ans++
		}
	}
	return ans
}

// 第二题
// leetcode704: 二分查找
func searchBinary(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 第四题
// leetcode35: 搜索插入位置
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

// 第五题
// leetcode217: 存在重复元素
func containsDuplicate(nums []int) bool {
	rec := append(sort.IntSlice{}, nums...)
	rec.Sort()
	for i := 0; i < len(rec)-1; i++ {
		if rec[i] == rec[i+1] {
			return false
		}
	}
	return true
}
