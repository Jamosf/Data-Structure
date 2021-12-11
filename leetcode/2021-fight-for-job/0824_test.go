// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

// tag-[单调栈]
// leetcode581: 最短无序连续子数组
// TODO
func findUnsortedSubarray(nums []int) int {
	left, right := -1, -1
	stack := make([]int, 0, len(nums))
	for i := range nums {
		for len(stack) != 0 && nums[stack[len(stack)-1]] >= nums[i] {
			if left == -1 {
				left = stack[len(stack)-1]
			} else {
				right = i
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	if right == -1 || left == -1 {
		return 0
	}
	return right - left + 1
}

func Test_findUnsortedSubarray(t *testing.T) {
	fmt.Println(findUnsortedSubarray([]int{1, 2, 2, 2, 3}))
}
