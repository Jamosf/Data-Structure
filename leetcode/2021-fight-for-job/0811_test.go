// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import "math"

// tag-[前缀和]
// 前缀和
// 第一题
// leetcode303: 区域和检索-数组不可变
type NumArray struct {
	sum []int
}

func ConstructorNumArray(nums []int) NumArray {
	numArray := NumArray{sum: make([]int, len(nums))}
	numArray.sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		numArray.sum[i] = numArray.sum[i-1] + nums[i]
	}
	return numArray
}

func (n *NumArray) SumRange(left int, right int) int {
	if left < 1 {
		return n.sum[right]
	}
	return n.sum[right] - n.sum[left-1]
}

// tag-[前缀和]
// 第二题
// leetcode1413: 逐步求和得到正数的最小值
func minStartValue(nums []int) int {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	minn := sum[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
		minn = min(minn, sum[i])
	}
	ans := 1 - minn
	if ans <= 0 {
		return 1
	}
	return ans
}

// tag-[前缀和]
// 第三题
// leetcode1480: 一维数组的动态和
func runningSum(nums []int) []int {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	return sum
}

// tag-[前缀和]
// 前缀和
// 第一题
// leetcode1732: 找到最高海拔
func largestAltitude(gain []int) int {
	maxn := math.MaxInt32
	sum := gain[0]
	for i := 1; i < len(gain); i++ {
		sum += gain[i]
		maxn = max(maxn, sum)
	}
	return maxn
}

// tag-[前缀和]
// 第三题
// leetcode 剑指offerII 012：左右两边子数组的和相等
func pivotIndex(nums []int) int {
	sum := make([]int, len(nums)+2)
	sum[0] = 0
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	sum[len(sum)-1] = sum[len(sum)-2]
	for i := 1; i < len(sum)-1; i++ {
		if sum[i-1] == sum[len(nums)-1]-sum[i] {
			return i - 1
		}
	}
	return -1
}
