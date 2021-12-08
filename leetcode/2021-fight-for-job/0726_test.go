// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

// tag-[数组]
// 第一题
// leetcode 剑指offer 30: 和为s的两个数字
func twoSum(nums []int, target int) []int {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return []int{target - v, v}
		}
		m[target-v] = struct{}{}
	}
	return nil
}
