// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

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
