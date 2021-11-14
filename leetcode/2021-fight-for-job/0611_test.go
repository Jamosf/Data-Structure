// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// leetcode713: 乘积小于k的子数组
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k < 1 {
		return 0
	}
	ans := 0
	m := 1
	left := 0
	for right, val := range nums {
		m *= val
		for m >= k {
			m /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}

func Test_num(t *testing.T) {
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
}
