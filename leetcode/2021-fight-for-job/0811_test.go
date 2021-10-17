// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import "math"

// 前缀和
// 第一题
type NumArray struct {
	sum []int
}

func Constructor1(nums []int) NumArray {
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

// 第二题
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

// 第三题
func runningSum(nums []int) []int {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	return sum
}

// 前缀和
// 第一题
func largestAltitude(gain []int) int {
	maxn := math.MaxInt32
	sum := gain[0]
	for i := 1; i < len(gain); i++ {
		sum += gain[i]
		maxn = max(maxn, sum)
	}
	return maxn
}

// 第二题
func isCovered(ranges [][]int, left int, right int) bool {
	//sort.Slice(ranges, func(x, y int) bool {
	//	return ranges[x][0] < ranges[y][0]
	//})
	//start, end := ranges[0][0], ranges[0][1]
	//for x := 1; x < len(ranges); x++ {
	//	if ranges[x][0]
	//}
	return false
}

// 第三题
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
